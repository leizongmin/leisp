'use strict';

register('build', function () {

  const pkg = 'github.com/leizongmin/leisp';
  const out = 'bin/leisp';

  const pkgParent = path.dirname(pkg);
  const gopath = path.resolve(os.tmpDir(), utils.randomString(10));
  env.GOPATH = gopath;
  mexec([
    `mkdir -p ${gopath}/src/${pkgParent}`,
    `ln -s ${pwd} ${gopath}/src/${pkg}`,
    `go build -o ${out} ${pkg}`,
    `rm -rf ${gopath}`,
  ]);

});
