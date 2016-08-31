'use strict';

// 自定义的参数
const pkg = 'github.com/leizongmin/leisp';
const out = 'bin/leisp';


register('build', function () {
  const pkgParent = path.dirname(pkg);
  const gopath = path.resolve(os.tmpDir(), utils.randomString(10));
  env.GOPATH = gopath;
  mexec([
    `mkdir -p ${gopath}/src/${pkgParent}`,
    `ln -s ${pwd} ${gopath}/src/${pkg}`,
    `go build -o ${out} ${pkg}`,
    `rm -rf ${gopath}`
  ]);
});

register('fetch', function () {
  if (argv.length < 1) return exit(1, 'run featch packages');
  const pkgParent = path.dirname(pkg);
  const gopath = path.resolve(os.tmpDir(), utils.randomString(10));
  env.GOPATH = gopath;
  if (!fs.existsSync(path.resolve(pwd, 'vendor'))) {
    exec(`mkdir ${pwd}/vendor`);
  }
  mexec([
    `mkdir -p ${gopath}`,
    `ln -s ${pwd}/vendor ${gopath}/src`,
    `go get ${argv.join(' ')}`,
    `rm -rf ${gopath}`,
  ]);
});
