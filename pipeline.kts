pipeline("dev"){
    parallel("web"){
        workspace("./web")
        cmd("yarn dev")
    }
    parallel("tailwind"){
        workspace("./web")
        cmd("npx tailwindcss -i./src/style.css -o./src/output.css")
    }
}
pipeline("api"){
    step("init"){
         workspace("./")
         cmd("hz new -module github.com/yanguiyuan/cloudspace  -idl idl/api/file.proto")
         cmd("go mod tidy")
    }
    step("clean"){
        workspace("./")
        cmd("Remove-Item -Path ./biz -Recurse ")
        cmd("Remove-Item -Path ./script -Recurse ")
        cmd("del ./.hz")
        cmd("del ./build.sh")
        cmd("del ./router.go")
        cmd("del ./router_gen.go")
        cmd("del ./main.go")
    }

    step("update"){
        cmd("hz update -I idl -idl idl/api/file.proto")
        cmd("hz update -I idl -idl idl/api/user.proto")
    }
    step("server"){
        cmd("go run .")
    }
}
pipeline("user"){
    step("init"){
        cmd("kitex -gen-path pkg -module github.com/yanguiyuan/cloudspace -service user  idl/user/user.thrift")
        movefile("main.go","./cmd/user/main.go")
        movefile("handler.go","./internal/user/handler/handler.go")
        movefile("kitex_info.yaml","./internal/user/kitex_info.yml")
        cmd("del build.sh")
        cmd("go run ./cmd/gen/user")
    }
    step("update"){
        cmd("kitex -gen-path pkg -module github.com/yanguiyuan/cloudspace idl/user/user.thrift")
        cmd("go run ./cmd/gen/user")
    }
}
pipeline("file"){
    step("update"){
        cmd("kitex -gen-path pkg -module github.com/yanguiyuan/cloudspace idl/cloudfile/cloudfile.thrift")
    }
}
pipeline("auth"){
    parallel("rpc"){
        cmd("go run ./cmd/user")
    }
    parallel("api"){
        cmd("go run .")
    }
    parallel("web"){
        workspace("./web")
        cmd("yarn dev")
    }
}
pipeline("server"){
    parallel("user"){
        cmd("go run ./cmd/user")
    }
    parallel("file"){
        cmd("go run ./cmd/cloudfile")
    }
    parallel("api"){
        cmd("go run ./cmd/api")
    }
}