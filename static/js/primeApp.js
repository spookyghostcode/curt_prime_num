var primeApp = angular.module("primeApp", []);

primeApp.controller("primeController", function ($scope, $http) {
    $scope.maxNum = 1000;
    $scope.primeNums = [];
    $scope.error = ""

    $scope.calcPrimes = function(maxNum) { 
        $scope.primeNums = []
        $scope.error = ""
        $http({
            url: "/primes",
            method: "GET",
            params: {max: maxNum}
        }).then(function successCallback(response) {
            $scope.primeNums = response.data.Success

        }, function errorCallback(response) {
            $scope.error = response.data.Error
        }); 
    }        
            
});
