require("colors");

var gulp = require("gulp"),
    nodemon = require("gulp-nodemon"),
    path = require("path"),
    spawn = require("child_process").spawn,
    nodemonInstance,
    env = {
      GOPATH: path.resolve(),
    };

gulp.task("run", function () {
    "use strict";
    if (nodemonInstance) {
        nodemonInstance.emit("restart");
    } else {
        startApp(env);
        return gulp.watch("./**/*.go", ["run"]);
    }
});

function startApp (env) {
    "use strict";
    nodemonInstance = nodemon({
        cwd: "src",
        watch: path.join("src", "__do_not_watch__"),
        script: "main.go",
        execMap: {
            "go": "go run"
        },
        env: env
    });
}
