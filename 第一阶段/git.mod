进入存放项目目录的文件夹目录
git init
git add +要上传的代码目录
git commit -m "add file"
git push -u origin main

常见错误：
error: failed to push some refs to 'github.com:Sion-L/golang.git'
hint: Updates were rejected because the remote contains work that you do
hint: not have locally. This is usually caused by another repository pushing
hint: to the same ref. You may want to first integrate the remote changes
hint: (e.g., 'git pull ...') before pushing again.
hint: See the 'Note about fast-forwards' in 'git push --help' for details.

解决：
先git pull 再上传
