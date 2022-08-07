package brackets

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBrackets_IsValid(t *testing.T) {
	{
		s := "([](){([])})"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 0, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "()[]}"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 5, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "{{[()]]"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 7, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "{{{[][][]"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 3, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "{*{{}"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 3, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "[[*"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 2, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "{*}"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 0, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "{{"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 2, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "{}"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 0, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := ""
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 0, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "}"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 1, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "*{}"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 0, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "{{{**[][][]"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 3, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "[]"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 0, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "{}[]"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 0, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "[()]"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 0, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "(())"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 0, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "{[]}()"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 0, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "{"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 1, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "{[}"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 3, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "foo(bar);"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 0, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "foo(bar[i);"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 10, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "[]([];"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 3, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "{[]"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 1, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "{{{"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 3, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "[]([]"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 3, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "{{{[][][]"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 3, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "{{{{{{{((()))}"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 6, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "{()}{"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 5, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "}()"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 1, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "()}()"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 3, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "}()"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 1, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "{[()]}}()"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 7, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "dasdsadsadas]]]"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 13, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "{}()"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 0, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "({}[(((())))])"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 0, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "()"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 0, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "({})"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 0, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "foo(bar({ <some initialization> })[i]);"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 0, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "([](){([])})"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 0, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "({["
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 3, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "({{}"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 2, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "()({}"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 3, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "[]([]"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 3, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "((({[]})"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 2, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "{}([]"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 3, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "(slkj, {lk[lve]} ,l)"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 0, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "(slkj{lk[lsj]}"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 1, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}

	{
		s := "{{[()]}"
		br := MustNewBracket()
		res := br.Validate(s)
		require.Equal(t, 1, res)
		if res > 0 {
			fmt.Println(res)
		} else {
			fmt.Println("Success")
		}
	}
}
