package main

import "fmt"

func main() {
	logScript(func() {
		awsRepo := setUpAwsRepo()

		//Use functions of aws repository
		ListBucketObjets(awsRepo)

		commonText := "factura001"
		filenameWithPrefix := fmt.Sprintf("/a2023/enero/%s.txt", commonText)

		//err := awsRepo.DownloadFile(filenameWithPrefix, fmt.Sprintf("../data/%s.txt", commonText))
		err := awsRepo.SaveFile(filenameWithPrefix, []byte(fmt.Sprintf("%s en formato texto bla bla...", commonText)))
		if err != nil {
			panic(err)
		}
		ListBucketObjets(awsRepo)
	})
}

func setUpAwsRepo() *AwsRepository {
	conf := GetConfig()
	return NewAwsRepo(&conf.Aws)
}
func logScript(f func()) {
	fmt.Println("Script starting ...")
	f()
	fmt.Println("Script finished successfully...")
}

func ListBucketObjets(awsRepo *AwsRepository) {
	fmt.Println("ListBucketObjets: ")
	objects := awsRepo.ListObjectsFromBucket()
	if sizeObj := len(objects); sizeObj == 0 {
		fmt.Println("------------------------------------")
		fmt.Println("No objects")
		fmt.Println("------------------------------------")
	} else {
		fmt.Println(fmt.Sprintf("Objects found %v", sizeObj))
	}

	for i, obj := range objects {
		fmt.Println("#Obj:         ", i)
		fmt.Println("Name:         ", *obj.Key)
		fmt.Println("Last modified:", *obj.LastModified)
		fmt.Println("Size:         ", *obj.Size)
		fmt.Println("Storage class:", *obj.StorageClass)
		fmt.Println("------------------------------------")
	}
	fmt.Println("- ListBucketObjets END -")
}
