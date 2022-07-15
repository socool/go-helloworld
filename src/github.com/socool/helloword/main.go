package main
func main() {
   println("Hello World")
   // VARIABLE DECLARATIONS
   println("TYPE 1")
   var a,b,c bool;
   println(a,":",b,":",c);

   println("TYPE 2")
   var d, e, f string;
   println(d,":",e,":",f);

   println("TYPE 3")
   var g, h, i bool = false, false, true;
   println(g,":",h,":",i);

   println("TYPE 4")
   var j, k, l = false, 10, "เรียนภาษา Go";
   println(j,":",k,":",l);

   printSchoolAddress();
   printSchoolAddressWithParam("ศรีราชา");

   schoolAddress := getSchoolAddress();
   println(schoolAddress);

   resultCode, resultAddress := getSchoolAddressWithParam(); 
   println(resultCode,":",resultAddress);

}
// FUNCTION DECLARATIONS
/**
BASIC FUNCTION
**/
func printSchoolAddress() {
   println("กรุงเทพ");
}

/**
WITH PARAMETER
**/
func printSchoolAddressWithParam(schoolAddress string) {
   println(schoolAddress);
}

/**
RETURN VALUE
**/
func getSchoolAddress() string {
   return "กรุงเทพ";
}


/**
RETURN VALUE
**/
func getSchoolAddressWithParam() (int, string) {
   code := 1993;
   address := "กรุงเทพ";
   return code, address;
}