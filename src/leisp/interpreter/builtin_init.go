// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package interpreter

var initLeispPrograms = `

(defn null? [a]
  (equal? (typeof a) "null")

`
