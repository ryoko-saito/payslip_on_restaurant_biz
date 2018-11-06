## 京都府版飲食関係給与明細計算機
<br><br>
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
