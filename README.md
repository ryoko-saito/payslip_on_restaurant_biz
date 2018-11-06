# 京都府の飲食関係給与明細計算機<br>
![サンプル画像](https://github.com/ryoko-saito/payslip_on_restaurant_biz/blob/master/img/payslip_calc.jpg)
<br><br>
京都府で飲食を営んでいて、京都府料理業組合連合会の保険に加入している方向けの給与明細を作成しました。<br>
※甲か乙か、従たる給与についての扶養控除などの申告書を提出しているかどうか、40歳以上の扶養人数などで保険料や所得税が変ってくるので、各項目は必ず入力して下さい。<br>
※賞与が発生する場合は、前月の給与（社会保険などの金額を控除したもの）を必ず入力して下さい。
<br><br>
## 使い方<br>
**Go言語のインストール(Windows機のみ)**<br>
[WindowsでGo言語 - saitodev.co](https://saitodev.co/article/1569)
<br><br>
**Qtのインストール**<br>
[Go言語でQtを扱ってみる on Windows - saitodev.co](https://saitodev.co/article/1956)<br>
[Go言語でQtを扱ってみる on Ubuntu - saitodev.co](https://saitodev.co/article/1949)
<br><br>
```
# 事前に税や保険の計算用のパッケージ取得する
go get github.com/ryoko-saito/tax_withholding/tax
cd サンプルコードのクローンを配置したいディレクトリ
git clone https://github.com/ryoko-saito/payslip_on_restaurant_biz.git
cd payslip_on_restaurant_biz
qtdeploy test desktop ./calc
```
## 注釈<br>
所得税の計算は、給与所得の源泉徴収税額表（平成30年分）から値を入力しました。<br>
賞与の計算は、賞与に対する源泉徴収税額の算出率の表（平成30年分）から値を入力しました。<br>
※前月中の給与等の金額がない場合や、前月中の給与等の金額が前月中の社会保険料等の金額以下である場合、
又はその賞与の金額（その金額から控除される社会保険料等の金額がある場合には、その控除後の金額）が前月中の給与等の
金額から前月中の社会保険料等の金額を控除した金額の10倍に相当する金額を超える場合には、別途月額表を使って税額を計算してください。<br>
健康保険料は、京都府料理業組合連合会（http://www.kyo-ryoinren.com/index.htm）の保険料をもとに算出しています。
前年度の住民税が非課税世帯の場合は、このアプリで健康保険料は算出できません。前年度課税世帯のみを対象としています。
　
