package main

func main() {

}

// func getInsuranceAmount(status insuranceStatus) int {
//   amount := 0
//   if !status.hasInsurance(){
//     amount = 1
//   } else {
//     if status.isTotaled(){
//       amount = 10000
//     } else {
//       if status.isDented(){
//         amount = 160
//         if status.isBigDent(){
//           amount = 270
//         }
//       } else {
//         amount = 0
//       }
//     }
//   }
//   return amount
// }

//THIS IS BETTER:

// func getInsuranceAmount(status insuranceStatus) int {
//   if !status.hasInsurance(){
//     return 1
//   }
//   if status.isTotaled(){
//     return 10000
//   }
//   if !status.isDented(){
//     return 0
//   }
//   if status.isBigDent(){
//     return 270
//   }
//   return 160
// }