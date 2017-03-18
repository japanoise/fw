# fw

fw converts ascii text to fullwidth text. It acts on its arguments; or stdin if
called with no arguments.

## Golf Challenge

There is an open challenge to implement this program in as few bytes as possible;
Aelita currently leads with [180 bytes of Perl.](https://gist.github.com/AelitaStyles/929189fbce048cb7aa8bb02d5b6982e1)
ZirconiumX holds second place at the moment with 
[their Python implementation.](https://gist.github.com/ZirconiumX/283a6cffba7b7e10ef62157563d11277)

Here's how the program should behave:

~~~
$ fw Hello World
Ｈｅｌｌｏ　Ｗｏｒｌｄ！

$ echo Hello World! | fw
Ｈｅｌｌｏ　Ｗｏｒｌｄ！

$ ascii
|00 nul|01 soh|02 stx|03 etx|04 eot|05 enq|06 ack|07 bel|
|08 bs |09 ht |0a nl |0b vt |0c np |0d cr |0e so |0f si |
|10 dle|11 dc1|12 dc2|13 dc3|14 dc4|15 nak|16 syn|17 etb|
|18 can|19 em |1a sub|1b esc|1c fs |1d gs |1e rs |1f us |
|20 sp |21  ! |22  " |23  # |24  $ |25  % |26  & |27  ' |
|28  ( |29  ) |2a  * |2b  + |2c  , |2d  - |2e  . |2f  / |
|30  0 |31  1 |32  2 |33  3 |34  4 |35  5 |36  6 |37  7 |
|38  8 |39  9 |3a  : |3b  ; |3c  < |3d  = |3e  > |3f  ? |
|40  @ |41  A |42  B |43  C |44  D |45  E |46  F |47  G |
|48  H |49  I |4a  J |4b  K |4c  L |4d  M |4e  N |4f  O |
|50  P |51  Q |52  R |53  S |54  T |55  U |56  V |57  W |
|58  X |59  Y |5a  Z |5b  [ |5c  \ |5d  ] |5e  ^ |5f  _ |
|60  ` |61  a |62  b |63  c |64  d |65  e |66  f |67  g |
|68  h |69  i |6a  j |6b  k |6c  l |6d  m |6e  n |6f  o |
|70  p |71  q |72  r |73  s |74  t |75  u |76  v |77  w |
|78  x |79  y |7a  z |7b  { |7c  | |7d  } |7e  ~ |7f del|


$ ascii | fw
｜００　ｎｕｌ｜０１　ｓｏｈ｜０２　ｓｔｘ｜０３　ｅｔｘ｜０４　ｅｏｔ｜０５　ｅｎｑ｜０６　ａｃｋ｜０７　ｂｅｌ｜
｜０８　ｂｓ　｜０９　ｈｔ　｜０ａ　ｎｌ　｜０ｂ　ｖｔ　｜０ｃ　ｎｐ　｜０ｄ　ｃｒ　｜０ｅ　ｓｏ　｜０ｆ　ｓｉ　｜
｜１０　ｄｌｅ｜１１　ｄｃ１｜１２　ｄｃ２｜１３　ｄｃ３｜１４　ｄｃ４｜１５　ｎａｋ｜１６　ｓｙｎ｜１７　ｅｔｂ｜
｜１８　ｃａｎ｜１９　ｅｍ　｜１ａ　ｓｕｂ｜１ｂ　ｅｓｃ｜１ｃ　ｆｓ　｜１ｄ　ｇｓ　｜１ｅ　ｒｓ　｜１ｆ　ｕｓ　｜
｜２０　ｓｐ　｜２１　　！　｜２２　　＂　｜２３　　＃　｜２４　　＄　｜２５　　％　｜２６　　＆　｜２７　　＇　｜
｜２８　　（　｜２９　　）　｜２ａ　　＊　｜２ｂ　　＋　｜２ｃ　　，　｜２ｄ　　－　｜２ｅ　　．　｜２ｆ　　／　｜
｜３０　　０　｜３１　　１　｜３２　　２　｜３３　　３　｜３４　　４　｜３５　　５　｜３６　　６　｜３７　　７　｜
｜３８　　８　｜３９　　９　｜３ａ　　：　｜３ｂ　　；　｜３ｃ　　＜　｜３ｄ　　＝　｜３ｅ　　＞　｜３ｆ　　？　｜
｜４０　　＠　｜４１　　Ａ　｜４２　　Ｂ　｜４３　　Ｃ　｜４４　　Ｄ　｜４５　　Ｅ　｜４６　　Ｆ　｜４７　　Ｇ　｜
｜４８　　Ｈ　｜４９　　Ｉ　｜４ａ　　Ｊ　｜４ｂ　　Ｋ　｜４ｃ　　Ｌ　｜４ｄ　　Ｍ　｜４ｅ　　Ｎ　｜４ｆ　　Ｏ　｜
｜５０　　Ｐ　｜５１　　Ｑ　｜５２　　Ｒ　｜５３　　Ｓ　｜５４　　Ｔ　｜５５　　Ｕ　｜５６　　Ｖ　｜５７　　Ｗ　｜
｜５８　　Ｘ　｜５９　　Ｙ　｜５ａ　　Ｚ　｜５ｂ　　［　｜５ｃ　　＼　｜５ｄ　　］　｜５ｅ　　＾　｜５ｆ　　＿　｜
｜６０　　｀　｜６１　　ａ　｜６２　　ｂ　｜６３　　ｃ　｜６４　　ｄ　｜６５　　ｅ　｜６６　　ｆ　｜６７　　ｇ　｜
｜６８　　ｈ　｜６９　　ｉ　｜６ａ　　ｊ　｜６ｂ　　ｋ　｜６ｃ　　ｌ　｜６ｄ　　ｍ　｜６ｅ　　ｎ　｜６ｆ　　ｏ　｜
｜７０　　ｐ　｜７１　　ｑ　｜７２　　ｒ　｜７３　　ｓ　｜７４　　ｔ　｜７５　　ｕ　｜７６　　ｖ　｜７７　　ｗ　｜
｜７８　　ｘ　｜７９　　ｙ　｜７ａ　　ｚ　｜７ｂ　　｛　｜７ｃ　　｜　｜７ｄ　　｝　｜７ｅ　　～　｜７ｆ　ｄｅｌ｜
~~~

All ASCII punctuation must be translated to its fullwidth representation; spaces should be converted
to the Unicode code point U+3000 (Ideographic space).
