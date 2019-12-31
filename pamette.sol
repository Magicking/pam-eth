pragma solidity >=0.4.22 <0.7.0;

contract Pamette {
	address public owner;

	mapping(bytes32 => address) public AllowedUser;

	constructor() public {
		owner = msg.sender;
	}

	function setUser(bytes32 hash, address user) public {
		require(msg.sender == owner, "Only owner can use this function");
		// TODO handle list
		if (AllowedUser[hash] != address(0)) {
			delete AllowedUser[hash];
		} else {
			AllowedUser[hash] = user;
		}
	}

	function isAuthorized(uint256 userId, string memory userName, string memory sharedPassword, uint256 blockNumber, uint8 v, bytes32 r, bytes32 s) public view
		returns (bool)
	{
		require(block.number - blockNumber <= 5, "Hash time invalid");
		bytes32 userHash = generateUserHash(userId, userName, sharedPassword);
		require(AllowedUser[userHash] != address(0), "User inexistant");
		bytes32 hashToSign = keccak256(abi.encode(userHash, block.number));

		bytes memory prefix = "\x19Ethereum Signed Message:\n13";
		bytes32 prefixedHash = keccak256(abi.encodePacked(prefix, hashToSign));
		require(AllowedUser[userHash] == ecrecover(prefixedHash, v, r, s), "User unpermitted");

		return true;
	}

	function generateUserHash(uint256 userId, string memory userName, string memory sharedPassword) public pure
	    returns (bytes32)
	{
		return keccak256(abi.encode(userName, userId, sharedPassword));
	}

	function generateOTP(uint256 userId, string memory userName, string memory sharedPassword) public view
		returns (bytes32 , uint)
	{
		bytes32 userHash = generateUserHash(userId, userName, sharedPassword);
		// Contatenate user + sharedPassword + currentBlockNumber
		bytes32 hashToSign = keccak256(abi.encode(userHash, block.number));

		return (hashToSign, block.number);
	}
}
