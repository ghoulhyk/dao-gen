以下是 `text/template` 包中的全面语法列表及其说明：

1. **插入变量或表达式**：使用 `{{ . }}` 插入变量或表达式的值。

    - 示例：`{{ .Name }}` 表示插入 `Name` 变量的值。

2. **管道操作符**：使用 `{{ .Field | Func1 | Func2 }}` 对变量应用管道操作符。

    - 示例：`{{ .Price | printf "%.2f" | upper }}` 表示对 `Price` 变量的值先应用 `printf` 函数，再应用 `upper` 函数。

3. **条件判断**：使用 `{{ if .Condition }} ... {{ else }} ... {{ end }}` 进行条件判断。

    - 示例：
      ```
      {{ if .HasPermission }}
          <p>Welcome, {{ .Username }}!</p>
      {{ else }}
          <p>Access denied.</p>
      {{ end }}
      ```

4. **循环遍历**：使用 `{{ range .List }} ... {{ end }}` 遍历一个列表。

    - 示例：
      ```
      <ul>
      {{ range .Items }}
          <li>{{ . }}</li>
      {{ end }}
      </ul>
      ```

5. **注释**：使用 `{{/* 注释内容 */}}` 对模板进行注释。

    - 示例：`{{/* This is a comment */}}`

6. **定义模板块**：使用 `{{ define "TemplateName" }} ... {{ end }}` 来定义一个模板块。

    - 示例：
      ```
      {{ define "Greeting" }}
          Hello, {{ .Name }}!
      {{ end }}
      ```

7. **引用模板块**：使用 `{{ template "TemplateName" . }}` 来引用另一个已定义的模板块。

    - 示例：`{{ template "Greeting" .User }}`

8. **定义变量**：使用 `{{ $var := .Value }}` 定义一个局部变量。

    - 示例：`{{ $name := .FirstName }}`

9. **执行函数**：使用 `{{ FuncName .Arg1 .Arg2 }}` 来调用自定义函数。

    - 示例：`{{ len .Items }}`

10. **原始字符串输出**：使用 `{{- ... -}}` 进行原始字符串输出，去除前后空白字符。

    - 示例：`{{- "   hello   " -}}`

11. **模板注入**：使用 `{{ block "BlockName" . }} ... {{ end }}` 在模板中定义可被覆盖的注入点。

    - 示例：
      ```
      {{ block "content" . }}
          <p>This is the default content.</p>
      {{ end }}
      ```

这些是常见的 `text/template` 语法和功能。还有其他一些高级用法和内置函数可供探索。在实际使用中，你可以根据具体需求选择并应用合适的语法来生成定制化的输出。

`text/template` 是 Go
语言的一个模板引擎，它提供了一系列内置函数来在模板中进行文本处理和逻辑控制。以下是一些常用的 `text/template` 内置函数：

1. `and`: 返回所有传入参数的逻辑与结果。
2. `or`: 返回所有传入参数的逻辑或结果。
3. `not`: 返回传入参数的逻辑非结果。
4. `eq`: 判断两个值是否相等。
5. `ne`: 判断两个值是否不相等。
6. `lt`: 判断一个值是否小于另一个值。
7. `le`: 判断一个值是否小于等于另一个值。
8. `gt`: 判断一个值是否大于另一个值。
9. `ge`: 判断一个值是否大于等于另一个值。
10. `len`: 返回传入参数的长度（对字符串、数组、切片、字典等类型有效）。
11. `index`: 根据索引返回数组、切片或字符串中的元素。
12. `slice`: 根据起始索引和结束索引返回切片或字符串的子序列。
13. `range`: 在循环中迭代数组、切片、字典或字符串。
14. `with`: 在一个新的作用域中执行模板块。
15. `printf`: 格式化字符串并返回结果。
16. `html`: 对字符串进行 HTML 转义。
17. `js`: 对字符串进行 JavaScript 转义。
18. `urlquery`: 对字符串进行 URL 转义。
19. `call`: 调用指定名称的自定义函数。

这仅是一些常见的内置函数，`text/template` 还提供了其他一些函数用于模板处理。你可以根据需要在模板中使用这些内置函数来完成文本处理和逻辑控制的任务。详细的内置函数列表和用法可以参考
Go 官方文档中 `text/template` 包的说明。

https://pkg.go.dev/text/template

http://masterminds.github.io/sprig/