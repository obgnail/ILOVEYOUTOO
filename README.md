## 缘起

著名的 2009 年 [贴吧爱情故事](https://tieba.baidu.com/p/529691897?pn=1)，此脚本帮你高效约会（笑

![iloveyoutoo](assets/iloveyoutoo.png)



## 使用

```go
func main() {
	rawtext := "ILOVEYOUTOO"
	result := encrypt(rawtext)
	fmt.Println(result) // ....-/.----/----./....-/....-/.----/---../.----/....-/.----/-..../...--/....-/.----/----./..---/-..../..---/..---/...--/--.../....-

	ciphertext := "....-/.----/----./....-/....-/.----/---../.----/....-/.----/-..../...--/....-/.----/----./..---/-..../..---/..---/...--/--.../....-/"
	result_ := decrypt(ciphertext)
	fmt.Println(result_) // ILOVEYOUTOO
}
```

