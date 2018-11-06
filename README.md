京都版飲食関係給与明細計算機
<br><br>
## 計算機を使ってみる<br>
**Go言語でQtを試すための環境構築**<br>
[Go言語でQtを扱ってみる on Windows - saitodev.co](https://saitodev.co/article/Go%E8%A8%80%E8%AA%9E%E3%81%A7Qt%E3%82%92%E6%89%B1%E3%81%A3%E3%81%A6%E3%81%BF%E3%82%8B_on_Windows)<br>
[Go言語でQtを扱ってみる on Ubuntu - saitodev.co](https://saitodev.co/article/Go%E8%A8%80%E8%AA%9E%E3%81%A7Qt%E3%82%92%E6%89%B1%E3%81%A3%E3%81%A6%E3%81%BF%E3%82%8B_on_Ubuntu)
<br><br>
```
# 事前に税や保険の計算用のパッケージ取得する
go get github.com/ryoko-saito/tax_withholding
cd サンプルコードのクローンを配置したいディレクトリ
git clone https://github.com/ryoko-saito/payslip_on_restaurant_biz.git
cd payslip_on_restaurant_biz
qtdeploy test desktop ./calc
```
