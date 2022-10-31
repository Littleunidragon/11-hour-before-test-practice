# 11-hour-before-test-practice
Name describes everything that going on actually
Копия чата с урока:

You8:08 AM
https://go.dev/play/p/E_7bGM-uB4I
Richard Shashalevich8:26 AM
https://go.dev/play/p/8UNZAbLKun3
Aleksandr Stukalov8:27 AM
https://goplay.tools/snippet/q0yujZVjx4L
You8:33 AM
https://go.dev/play/p/ozyw5DEWpUU
Илья8:37 AM
https://go.dev/play/p/7q6szhwSk10
Richard Shashalevich8:45 AM
https://go.dev/play/p/DYw7Vb6LZ5k
Aleksandr Stukalov8:52 AM
https://goplay.tools/snippet/ty7YbeEOvdH
Jaroslavs Samcuks9:08 AM
N ламп расположены в ряд. Каждая лампа может светить красным, жёлтым или зелёным цветом. Сначала они все светят красным. Задача состоит в том, чтобы все лампы светились зелёным. При этом разрешается переключать цвет ламп по следующим правилам:
- за один раз разрешается переключать только одну лампу;
- лампу можно переключить, если все лампы слева светятся одинаково, и их цвет отличается от цвета переключаемой лапы; новый цвет лампы также должен отличаться от цвета всех ламп, стоящих левее.
Jaroslavs Samcuks9:12 AM
https://goplay.tools/snippet/Y4Ge1zvAohr
Andrej Ceremnih9:26 AM
https://goplay.tools/snippet/eRBoAYQ6ohp

Anastasija Voropajeva:
10.06  - https://goplay.tools/snippet/lXRsx2fDao7

29.09 - https://goplay.tools/snippet/v6CvRaepvWS

10.10 - https://goplay.tools/snippet/l3WOeRvExO6

Jaroslavs Samčuks:
func split(n int, tail []int, f func([]int)) {
  if n == 0 { f(tail); return }
  for i := 1; i <= n; i++ {
    if len(tail) == 0 || tail[len(tail)-1] <= i {
      split(n-i, append(tail, i), f)
    }
  }
}

я ожидаю примерно такое решение. и тот факт, что ты встроила внутренний if в рекурсию - оч красиво

Anastasija Voropajeva:
У меня эта ссылка ведёт на filterDigits 🤔
