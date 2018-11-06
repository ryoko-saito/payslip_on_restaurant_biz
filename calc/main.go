package main

import (
	"os"
	"strconv"

	"github.com/ryoko-saito/tax_withholding/tax"
	"github.com/therecipe/qt/widgets"
)

//packageスコープ内に変数を定義しておくと、mainの中、外でも変数を使うことができる
var salalyInput *widgets.QLineEdit
var bonusInput *widgets.QLineEdit
var dayInput *widgets.QLineEdit
var trafficInput *widgets.QLineEdit
var traffictotalInput *widgets.QLineEdit
var healthinsuranceInput *widgets.QLineEdit
var paymenttotalInput *widgets.QLineEdit
var employmentinshulanceInput *widgets.QLineEdit
var careinsuranceInput *widgets.QLineEdit
var pensionInput *widgets.QLineEdit
var socialinsulancesumInput *widgets.QLineEdit
var incometaxInput *widgets.QLineEdit
var deductionInput *widgets.QLineEdit
var supportInput *widgets.QLineEdit
var ageInput *widgets.QLineEdit
var kouorotuBox *widgets.QComboBox
var beforeincomeInput *widgets.QLineEdit
var followBox *widgets.QComboBox
var netpaymentsInput *widgets.QLineEdit

//var municipaltaxInput *widgets.QLineEdit
var lastInput *widgets.QLineEdit

func main() {

	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(400, 250)
	window.SetWindowTitle("Salaly calculator")

	gridLayout := widgets.NewQGridLayout2()
	widget := widgets.NewQWidget(nil, 0)

	nameLabel := widgets.NewQLabel2("氏名", nil, 0)
	gridLayout.AddWidget(nameLabel, 0, 0, 0)

	nameInput := widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(nameInput, 1, 0, 0)

	ageLabel := widgets.NewQLabel2("40歳以上64歳未満の扶養人数", nil, 0)
	gridLayout.AddWidget(ageLabel, 0, 2, 0)

	ageInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(ageInput, 1, 2, 0)

	kouorotuLabel := widgets.NewQLabel2("甲または乙の選択", nil, 0)
	gridLayout.AddWidget(kouorotuLabel, 0, 4, 0)

	kouorotuBox = widgets.NewQComboBox(nil)
	kouorotuBox.AddItems([]string{"甲", "乙"})

	gridLayout.AddWidget(kouorotuBox, 1, 4, 0)

	followLabel := widgets.NewQLabel2("従たる給与についての扶養控除申告書の提出", nil, 0)
	gridLayout.AddWidget(followLabel, 0, 5, 0)

	followBox = widgets.NewQComboBox(nil)
	followBox.AddItems([]string{"有り", "無し"})

	gridLayout.AddWidget(followBox, 1, 5, 0)

	supportLabel := widgets.NewQLabel2("扶養人数", nil, 0)
	gridLayout.AddWidget(supportLabel, 2, 4, 0)

	supportInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(supportInput, 3, 4, 0)

	beforeincomeLabel := widgets.NewQLabel2("2年前所得", nil, 0)
	gridLayout.AddWidget(beforeincomeLabel, 0, 6, 0)

	beforeincomeInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(beforeincomeInput, 1, 6, 0)

	salalyLabel := widgets.NewQLabel2("給与", nil, 0)
	gridLayout.AddWidget(salalyLabel, 2, 0, 0)

	salalyInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(salalyInput, 3, 0, 0)

	bonusLabel := widgets.NewQLabel2("賞与", nil, 0)
	gridLayout.AddWidget(bonusLabel, 2, 2, 0)

	bonusInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(bonusInput, 3, 2, 0)

	dayLabel := widgets.NewQLabel2("日数", nil, 0)
	gridLayout.AddWidget(dayLabel, 4, 0, 0)

	trafficfeeLabel := widgets.NewQLabel2("交通費", nil, 0)
	gridLayout.AddWidget(trafficfeeLabel, 4, 2, 0)

	traffictotleLabel := widgets.NewQLabel2("交通費合計", nil, 0)
	gridLayout.AddWidget(traffictotleLabel, 4, 4, 0)

	dayInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(dayInput, 5, 0, 0)

	kakeruLabel := widgets.NewQLabel2("×", nil, 0)
	gridLayout.AddWidget(kakeruLabel, 5, 1, 0)

	trafficInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(trafficInput, 5, 2, 0)

	ikoruLabel := widgets.NewQLabel2("=", nil, 0)
	gridLayout.AddWidget(ikoruLabel, 5, 3, 0)

	traffictotalInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(traffictotalInput, 5, 4, 0)

	paymenttotalLabel := widgets.NewQLabel2("支給合計", nil, 0)
	gridLayout.AddWidget(paymenttotalLabel, 4, 6, 0)

	paymenttotalInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(paymenttotalInput, 5, 6, 0)

	healthinsuranceLabel := widgets.NewQLabel2("健康保険", nil, 0)
	gridLayout.AddWidget(healthinsuranceLabel, 6, 0, 0)

	healthinsuranceInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(healthinsuranceInput, 7, 0, 0)

	careinsuranceLabel := widgets.NewQLabel2("介護保険", nil, 0)
	gridLayout.AddWidget(careinsuranceLabel, 6, 2, 0)

	careinsuranceInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(careinsuranceInput, 7, 2, 0)

	//	pensionLabel := widgets.NewQLabel2("厚生年金", nil, 0)
	//gridLayout.AddWidget(pensionLabel, 6, 4, 0)

	//pensionInput = widgets.NewQLineEdit(nil)
	//gridLayout.AddWidget(pensionInput, 7, 4, 0)

	lastLabel := widgets.NewQLabel2("前月給与", nil, 0)
	gridLayout.AddWidget(lastLabel, 2, 6, 0)

	lastInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(lastInput, 3, 6, 0)

	employmentinshulanceLabel := widgets.NewQLabel2("雇用保険", nil, 0)
	gridLayout.AddWidget(employmentinshulanceLabel, 6, 4, 0)

	employmentinshulanceInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(employmentinshulanceInput, 7, 4, 0)

	incometaxLabel := widgets.NewQLabel2("所得税", nil, 0)
	gridLayout.AddWidget(incometaxLabel, 8, 0, 0)

	incometaxInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(incometaxInput, 9, 0, 0)

	//municipaltaxLabel := widgets.NewQLabel2("住民税", nil, 0)
	//gridLayout.AddWidget(municipaltaxLabel, 10, 2, 0)

	//municipaltaxInput = widgets.NewQLineEdit(nil)
	//gridLayout.AddWidget(municipaltaxInput, 11, 2, 0)

	socialinsulancesumLabel := widgets.NewQLabel2("社会保険合計", nil, 0)
	gridLayout.AddWidget(socialinsulancesumLabel, 6, 6, 0)

	socialinsulancesumInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(socialinsulancesumInput, 7, 6, 0)

	deductiontotalLabel := widgets.NewQLabel2("控除合計", nil, 0)
	gridLayout.AddWidget(deductiontotalLabel, 8, 6, 0)

	deductionInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(deductionInput, 9, 6, 0)

	netpaymentsLabel := widgets.NewQLabel2("差し引き支給合計", nil, 0)
	gridLayout.AddWidget(netpaymentsLabel, 10, 6, 0)

	netpaymentsInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(netpaymentsInput, 11, 6, 0)

	beforeincomeInput.ConnectEditingFinished(calcAndInsert)
	supportInput.ConnectEditingFinished(calcAndInsert)
	ageInput.ConnectEditingFinished(calcAndInsert)
	dayInput.ConnectEditingFinished(calcAndInsert)
	trafficInput.ConnectEditingFinished(calcAndInsert)
	salalyInput.ConnectEditingFinished(calcAndInsert)
	bonusInput.ConnectEditingFinished(calcAndInsert)
	traffictotalInput.ConnectEditingFinished(calcAndInsert)
	paymenttotalInput.ConnectEditingFinished(calcAndInsert)
	careinsuranceInput.ConnectEditingFinished(calcAndInsert)
	socialinsulancesumInput.ConnectEditingFinished(calcAndInsert)
	healthinsuranceInput.ConnectEditingFinished(calcAndInsert)
	//pensionInput.ConnectEditingFinished(calcAndInsert)
	//deductionInput.ConnectEditingFinished(calcAndInsert)
	deductionInput.ConnectEditingFinished(calcAndInsert)
	lastInput.ConnectEditingFinished(calcAndInsert)
	netpaymentsInput.ConnectEditingFinished(calcAndInsert)

	//box系のイベント
	followBox.ConnectCurrentIndexChanged2(func(text string) {
		calcAndInsert()
	})

	kouorotuBox.ConnectCurrentIndexChanged2(func(text string) {
		calcAndInsert()
	})

	widget.SetLayout(gridLayout)
	window.SetCentralWidget(widget)

	window.Show()
	app.Exec()
}

//各Inputの計算
func calcAndInsert() {
	dayV := convert(dayInput)
	trafficV := convert(trafficInput)
	res := dayV * trafficV
	insert(traffictotalInput, res)

	bonusAmount := convert(bonusInput)

	//支給合計
	salalyAmount := convert(salalyInput)
	trafficTotalAmount := convert(traffictotalInput)
	paymentTotalAmount := salalyAmount + trafficTotalAmount + bonusAmount
	insert(paymenttotalInput, paymentTotalAmount)

	//　雇用保険の金額

	employmentinshulanceAmount := paymentTotalAmount * 3 / 1000
	insert(employmentinshulanceInput, employmentinshulanceAmount)

	//扶養人数
	supportnumber := convert(supportInput)
	var healthinsuranceAmout int

	//健康保険の金額 前年度所得が入力されていることが処理の条件
	if len(beforeincomeInput.Text()) > 0 {
		beforeincomeAmout := convert(beforeincomeInput)

		//第二引数の計算は扶養人数 + 1にする
		healthinsuranceAmout = tax.CalcHealthInsurance(beforeincomeAmout, (supportnumber + 1))
		insert(healthinsuranceInput, healthinsuranceAmout)
	}

	//介護保険の計算
	carenumber := convert(ageInput)
	careinsuranceAmount := tax.CalcCareInsurance(carenumber)
	insert(careinsuranceInput, careinsuranceAmount)

	//厚生年金の計算
	//pensionAmout := convert(pensionInput)

	//社会保険の合計
	socialinsulancesumAmount := healthinsuranceAmout + careinsuranceAmount
	insert(socialinsulancesumInput, socialinsulancesumAmount)

	//所得（給与ー社会保険合計）
	incomeAmout := salalyAmount - socialinsulancesumAmount

	kouorotuIndex := kouorotuBox.CurrentIndex()

	//従たる給与についての扶養控除申告書の提出

	followIndex := followBox.CurrentIndex()

	//所得税の計算
	incometaxAmount := tax.CalcTax(incomeAmout, kouorotuIndex, supportnumber, followIndex)
	lastAmount := convert(lastInput)
	//賞与がある場合
	if lastAmount > 0 && bonusAmount > 0 {
		incometaxAmount += tax.CalcBonusTax(lastAmount, kouorotuIndex, supportnumber, followIndex)
	}
	insert(incometaxInput, incometaxAmount)

	//控除合計の計算
	deductionAmount := incometaxAmount
	insert(deductionInput, deductionAmount)

	//差し引き支給額の計算
	netpaymentsAmount := paymentTotalAmount - socialinsulancesumAmount - deductionAmount
	insert(netpaymentsInput, netpaymentsAmount)
}

func convert(e *widgets.QLineEdit) int {
	str := e.Text()
	if len(str) == 0 {
		return 0
	}

	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return i
}

func insert(e *widgets.QLineEdit, i int) {
	h := strconv.Itoa(i)
	e.Clear()
	e.Insert(h)
}
