package main

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"os"
	"time"
)

type Letter struct {
	Subject string
	Date time.Time
	Period int
	Student string
	Class string
	Pass bool
}

func LetterGen(letter Letter){
	begin := time.Now()
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	//m.SetBorder(true)

	m.Row(20, func() {
		m.Col(8, func() {
			m.Text(fmt.Sprintf("%s Department", letter.Student), props.Text{
				Top:   6,
				Align: consts.Left,
				Size:  20,
				Style: consts.Bold,
			})
		})
		m.Col(4, func() {
			_ = m.FileImage("internal/assets/images/tbs-logo.png", props.Rect{
				Percent: 90,
				Left:  28,
			})
		})
	})

	m.Row(20, func(){
		m.Col(12, func() {
			m.Text("Recuperação Paralela Results/Follow up", props.Text{
				Top:   6,
				Align: consts.Left,
				Size:  20,
				Style: consts.Bold,
			})
		})
	})

	m.Row(15, func(){
		m.Col(4, func() {
			m.Text(fmt.Sprintf("Date: %s", letter.Date.Format(time.RFC822)), props.Text{
				Top:   3,
				Align: consts.Left,
				Size:  15,
			})
		})
		m.Col(4, func() {
			m.Text(fmt.Sprintf("Marking Period: %d", letter.Period), props.Text{
				Top:   3,
				Align: consts.Left,
				Size:  15,
			})
		})
		m.Col(4, func() {
			m.Text(fmt.Sprintf("School Year: %d", letter.Date.Year()), props.Text{
				Top:   3,
				Align: consts.Left,
				Size:  15,
			})
		})
	})

	m.Row(15, func(){
		m.Col(6, func() {
			m.Text(fmt.Sprintf("Student: %s", letter.Student), props.Text{
				Top:   3,
				Align: consts.Left,
				Size:  15,
			})
		})
		m.Col(6, func() {
			m.Text(fmt.Sprintf("Class: %s", letter.Class), props.Text{
				Top:   3,
				Align: consts.Left,
				Size:  15,
			})
		})
	})

	m.Row(15, func(){
		m.Col(12, func() {
			m.Text("Dear parents/guardians,", props.Text{
				Top:   3,
				Align: consts.Left,
				Size:  15,
			})
		})
	})



	if letter.Pass {
		m.Row(20, func(){
			m.Col(12, func() {
				m.Text(fmt.Sprintf("The purpose of this communication is to inform that %s has met the minimum requirements to pass %s this marking period.", letter.Student, letter.Subject), props.Text{
					Top:   3,
					Align: consts.Left,
					Size:  15,
				})
			})
		})

		m.Row(20, func(){
			m.Col(12, func() {
				m.Text(fmt.Sprintf("According to the Recuperação Paralela requirements, %s grade on the report will be changed to a passing grade 5.", letter.Student), props.Text{
					Top:   3,
					Align: consts.Left,
					Size:  15,
				})
			})
		})
	}else{
		m.Row(20, func(){
			m.Col(12, func() {
				m.Text(fmt.Sprintf("The purpose of this communication is to inform that %s has not met the minimum requirements to pass %s this marking period.", letter.Student, letter.Subject), props.Text{
					Top:   3,
					Align: consts.Left,
					Size:  15,
				})
			})
		})

		m.Row(20, func(){
			m.Col(12, func() {
				m.Text(fmt.Sprintf("According to the Recuperação Paralela requirements, %s grade on the report will remain the same.", letter.Student), props.Text{
					Top:   3,
					Align: consts.Left,
					Size:  15,
				})
			})
		})
	}




	m.Row(10, func(){
		m.Col(12, func() {
			m.Text("If you have any questions, please do not hesitate in contacting me.", props.Text{
				Top:   3,
				Align: consts.Left,
				Size:  15,
			})
		})
	})

	m.Row(15, func(){
		m.Col(12, func() {
			m.Text("Thank you for your attention and support.", props.Text{
				Top:   6,
				Align: consts.Left,
				Size:  15,
			})
		})
	})

	m.Row(25, func() {
		m.Col(6, func() {
			m.Signature("Teacher")
		})
		m.Col(6, func() {
			m.Signature("Head of Department")
		})
	})

	m.Row(15, func() {
		m.ColSpace(12)
	})
	m.Line(1)

	m.Row(15, func(){
		m.Col(4, func() {
			m.Text("Date: ", props.Text{
				Top:   6,
				Align: consts.Left,
				Size:  15,
			})
		})
		m.Col(4, func() {
			m.Text(fmt.Sprintf("Marking Period: %d", letter.Period), props.Text{
				Top:   6,
				Align: consts.Left,
				Size:  15,
			})
		})
		m.Col(4, func() {
			m.Text(fmt.Sprintf("School Year: %d", letter.Date.Year()), props.Text{
				Top:   6,
				Align: consts.Left,
				Size:  15,
			})
		})
	})

	m.Row(15, func(){
		m.Col(6, func() {
			m.Text(fmt.Sprintf("Student: %s", letter.Student), props.Text{
				Top:   3,
				Align: consts.Left,
				Size:  15,
			})
		})
		m.Col(6, func() {
			m.Text(fmt.Sprintf("Class: %s", letter.Class), props.Text{
				Top:   3,
				Align: consts.Left,
				Size:  15,
			})
		})
	})

	m.Row(15, func(){
		m.Col(12, func() {
			m.Text(fmt.Sprintf("I have received the %s Recuperação Paralela Results/Follow up.", letter.Subject), props.Text{
				Top:   3,
				Align: consts.Left,
				Size:  15,
			})
		})
	})

	m.Row(25, func() {
		m.Col(6, func() {
			m.Signature("Student")
		})
		m.Col(6, func() {
			m.Signature("Parent")
		})
	})

	err := m.OutputFileAndClose("certificate.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}