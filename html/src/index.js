require('bootstrap')
require('angular')
const ethUtil = require('ethereumjs-util')

function extractParams() {
    let queryString = window.location.search.slice(1)
    let params = window.location.hash.split('#')[1]
    let arr = params.split('&')
    let ret = {}
    for (var i = 0; i < arr.length; i++) {
      var a = arr[i].split('=')
      ret[a[0]] = a[1]
    }
    return ret
}

angular.module('pamette',[]).controller('cardController', function($scope) {
	let ethereum
	let ret = extractParams()
	let setStatus = (timeLeft, buttonTxt, pgStatus, txtStatus) => {
		$scope.timeLeft = timeLeft
		$scope.buttonTxt = buttonTxt
		$scope.pgStatus = pgStatus
		$scope.txtStatus = txtStatus
		$scope.$apply()
	}
	if (typeof window.ethereum !== 'undefined') {
		ethereum = window.ethereum
		ethereum.send('eth_accounts')
		.then((ret) => {
			console.log(ret)
			if (ret.result.length == 0) {
				setStatus(0, 'Connect', 'bg-warning', 'Please connect on Goërli')
			} else {
				$scope.user = 'Personne'
				$scope.machineName = 'Emergence'
				setStatus(100, 'Authenticate', '', '')
			}
		})
		.catch(err => {
			// In the future, maybe in 2020, this will return a 4100 error if
			// the user has yet to connect
			if (err.code === 4100) { // EIP 1193 unauthorized error
				setStatus(0, 'Connect', 'bg-warning', 'Please connect on Goërli')
			} else {
				setStatus(0, '', 'bg-danger', err)
				console.error(err)
			}
		})
	} else {
		setStatus(0, '', 'bg-danger', 'Please install Metamask')
	}
	$scope.signHash = function () {
		if (!(typeof web3 !== 'undefined')) {
			console.log("No web3 provider present!")
			return
		}
		let from = web3.eth.accounts[0]
		if (!from) {
			if (typeof ethereum !== 'undefined') {
				ethereum.enable()
				.catch(console.error)
			}
			return
		}

		let msgBuf = Buffer.from(ret['token'], 'hex')
        var msg = ethUtil.bufferToHex(msgBuf)
		var params = [msg, from]
		var method = 'personal_sign'
		web3.currentProvider.sendAsync({
			method,
			params,
			from,
		}, function (err, result) {
			if (err) return console.error(err)
			if (result.error) return console.error(result.error)
			let sig = result.result
			console.log('Hash signature:', sig)
			$scope.txtStatus = sig
			$scope.$apply()
		})
	}
})
