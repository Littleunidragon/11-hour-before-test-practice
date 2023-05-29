# 11-hour-before-test-practice
Name describes everything that going on actually
Копия чата с урока:

https://go.dev/play/p/DYw7Vb6LZ5k
https://goplay.tools/snippet/ty7YbeEOvdH

N ламп расположены в ряд. Каждая лампа может светить красным, жёлтым или зелёным цветом. Сначала они все светят красным. Задача состоит в том, чтобы все лампы светились зелёным. При этом разрешается переключать цвет ламп по следующим правилам:
- за один раз разрешается переключать только одну лампу;
- лампу можно переключить, если все лампы слева светятся одинаково, и их цвет отличается от цвета переключаемой лапы; новый цвет лампы также должен отличаться от цвета всех ламп, стоящих левее.

https://goplay.tools/snippet/Y4Ge1zvAohr
https://goplay.tools/snippet/eRBoAYQ6ohp

10.06  - https://goplay.tools/snippet/lXRsx2fDao7
29.09 - https://goplay.tools/snippet/v6CvRaepvWS
10.10 - https://goplay.tools/snippet/l3WOeRvExO6

func split(n int, tail []int, f func([]int)) {
  if n == 0 { f(tail); return }
  for i := 1; i <= n; i++ {
    if len(tail) == 0 || tail[len(tail)-1] <= i {
      split(n-i, append(tail, i), f)
    }
  }
}
