package main

import "fmt"

func main() {
	fmt.Println("Script starting ...")
	conf := GetConfig()
	awsRepo := NewAwsRepo(&conf.Aws)
	ListBucketObjets(awsRepo)

	commonText := "factura001"
	filenameWithPrefix := fmt.Sprintf("/a2023/enero/%s.txt", commonText)

	err := awsRepo.DownloadFile(filenameWithPrefix, fmt.Sprintf("C:\\Users\\cassa\\GitRepos\\aws-s3-example\\data\\%s.txt", commonText))
	//err := awsRepo.SaveFile(filenameWithPrefix, []byte(fmt.Sprintf("%s en formato texto bla bla...", commonText)))
	if err != nil {
		panic(err)
	}

	ListBucketObjets(awsRepo)
	fmt.Println("Script finished successfully...")
}

func ListBucketObjets(awsRepo *AwsRepository) {
	objects := awsRepo.ListObjectsFromBucket()
	if len(objects) == 0 {
		fmt.Println("------------------------------------")
		fmt.Println("No objects")
		fmt.Println("------------------------------------")
	}
	for i, obj := range objects {
		fmt.Println("#Obj:         ", i)
		fmt.Println("Name:         ", *obj.Key)
		fmt.Println("Last modified:", *obj.LastModified)
		fmt.Println("Size:         ", *obj.Size)
		fmt.Println("Storage class:", *obj.StorageClass)
		fmt.Println("------------------------------------")
	}
}
