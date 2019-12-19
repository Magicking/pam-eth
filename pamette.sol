pragma solidity >=0.4.22 <0.7.0;

contract Pamette {
	address public owner;

	mapping(uint256 => bool) public AllowUser
	constructor() {
		owner = msg.sender;
	}

	
	function setUser(uint256 hash) void public {
		require(msg.sender == owner, "Only owner can use this function");
		if (AllowedUser[hash]) {
			delete AllowedUser[hash];
		} else {
			AllowedUser[hash] = true;
		}
	}
	
	function isAuthorized(bytes memory data, bytes memory signature) public view
		returns (bool)
	{
		return true;
	}
	
	function generateOTP() public view
		returns (string memory)
	{
		return "password";
	}
}
