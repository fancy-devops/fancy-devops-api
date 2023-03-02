package controller

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"

	"gitlab.chad122.top/fancy-devops/fancy-devops-api/model/codes"
	"gitlab.chad122.top/fancy-devops/fancy-devops-api/model/requests"
	"gitlab.chad122.top/fancy-devops/fancy-devops-api/pkg/setting"
	"gitlab.chad122.top/fancy-devops/fancy-devops-api/pkg/util"
)

type Git struct {
}

func NewGit() *Git {
	return &Git{}
}

// @Tags Git
// @Summary Git钩子
// @Accept json
// @Produce json
// @Param "" body requests.HookReq true "Git钩子"
// @Param token header string true "token"
// @Success 200 {object} responses.Result "成功"
// @Failure 400 {object} responses.Result "错误的请求"
// @Failure 403 {object} responses.Result "无权限"
// @Failure 500 {object} responses.Result "系统异常"
// @Router /api/git/hook [POST]
func (obj *Git) Hook(c *gin.Context) {
	token := c.GetHeader("X-Gitlab-Token")
	if token != setting.GitSecret {
		util.NewGinRes().ErrorResponse(c, codes.Common_Forbidden)
		return
	}

	var req requests.HookReq
	err := c.BindJSON(&req)
	if err != nil {
		util.NewGinRes().ErrorResponse2(c, codes.Common_BadRequest, "参数错误")
		return
	}

	// 判断仓库 & 分支是否配置了自动CI/CD
	if req.Repository.HttpUrl != "http://gitlab.chad122.top/fancy-devops/fancy-devops-api.git" {
		util.NewGinRes().ErrorResponse(c, codes.Common_Forbidden)
		return
	}

	// clone & checkout
	repPath := setting.FileBasePath + setting.RepositoryPath + "/" + req.Repository.Name
	_ = os.RemoveAll(repPath)
	rep, err := git.PlainClone(repPath, false, &git.CloneOptions{
		URL:      req.Repository.HttpUrl,
		Progress: os.Stdout,
		Auth: &http.BasicAuth{
			Username: "fancy-devops",
			Password: setting.GitAccessToken,
		},
	})
	if err != nil {
		util.NewGinRes().ErrorResponse2(c, codes.Common_Error, "clone 失败")
		return
	}

	wt, err := rep.Worktree()
	if err != nil {
		util.NewGinRes().ErrorResponse2(c, codes.Common_Error, "Worktree 失败")
		return
	}
	err = wt.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(req.Ref),
	})
	if err != nil {
		util.NewGinRes().ErrorResponse2(c, codes.Common_Error, "checkout 失败"+req.Ref)
		return
	}

	// build
	ts := time.Now().UnixMilli()
	imgName := fmt.Sprintf("fancy-devops-api:v%d", ts)
	buildCmd := exec.Command("docker", "build", "-t", imgName, ".")
	buildCmd.Dir = repPath

	var stdout io.ReadCloser
	if stdout, err = buildCmd.StdoutPipe(); err != nil {
		log.Fatal(err)
	}

	defer stdout.Close()
	if err := buildCmd.Start(); err != nil {
		log.Fatal(err)
	}

	if opBytes, err := ioutil.ReadAll(stdout); err != nil {
		log.Fatal(err)
	} else {
		log.Print(string(opBytes))
	}

	// deploy
	//loadCmd := exec.Command("minikube", "image", "load", imgName)

	util.NewGinRes().SuccessResponse(c, "")
}
