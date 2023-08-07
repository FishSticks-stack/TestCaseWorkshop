# How to Test your Code  ![Static Badge](https://img.shields.io/badge/GO-white?logo=Go) ![Static Badge](https://img.shields.io/badge/Visual%20Studio%20Code-white?logo=visualstudiocode&logoColor=blue)

*Workshop from <a href="https://ghw.mlh.io/events/data-week">MLH's GHW - Data Week</a>*

Using <a href="https://go.dev/doc/tutorial/getting-started">Go</a> we made our program to mimick payment functionality, including the order struct, payment processor, inventory managment, reduce stock. 

Adding in a file with unit test cases for any possible failures the program could encounter during its execution. 

Such as: 
* invalid orders
* invalid order quantity
* invalid inventory reduction error

Setup
--
To run this project make sure to have Golang installed, click here to <a href="https://go.dev/dl/">Download</a>.

Inside the Visual Studio Code terminal enter:
```
% go mod init TestCaseWorkshop
% go mod tidy
% go build
% ./TestCaseWorkshop
```

**Remember to cd into the project before entering commands*
<img width="1385" alt="image" src="https://github.com/FishSticks-stack/TestCaseWorkshop/assets/70287581/bd6cf20d-2106-4f8a-bdd3-e8c99cf24b6f">


If you need more help, see the <a href="https://go.dev/doc/tutorial/getting-started">Getting Started</a> guide.

Normal execution
---
Normal execution of the program should look like this.

<img width="580" alt="Screenshot 2023-08-06 at 7 47 44 PM" src="https://github.com/FishSticks-stack/TestCaseWorkshop/assets/70287581/08caba9c-e622-46e4-a7a8-efa0a20acd77">


  *Order processed successfully!*

Test Cases
---
Running all three of these test cases results in all passing cases.

* Valid order ✅
* Invalid order quantity ✅
* Invalid reduction error ✅
  
<img width="1170" alt="image" src="https://github.com/FishSticks-stack/TestCaseWorkshop/assets/70287581/caf110f0-e7e1-4515-bc37-5647c62e291a">


Now we know the code has a valid order, quantity, and the reduction is executed correctly through these test cases.






