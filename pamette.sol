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

	function isAuthorized(string memory userName, uint256 userId, uint256 machineId, bytes32 sharedPassword, uint256 blockNumber, uint8 v, bytes32 r, bytes32 s) public view
		returns (bool)
	{
		require(block.number - blockNumber <= 5, "Hash time invalid");
		bytes32 userHash = generateUserHash(userName, userId, machineId, sharedPassword);
		require(AllowedUser[userHash] != address(0), "User inexistant");
		bytes32 hashToSign = keccak256(abi.encode(userHash, block.number));

		bytes memory prefix = "\x19Ethereum Signed Message:\n13";
		bytes32 prefixedHash = keccak256(abi.encodePacked(prefix, hashToSign));
		require(AllowedUser[userHash] == ecrecover(prefixedHash, v, r, s), "User unpermitted");

		return true;
	}

	function generateUserHash(string memory userName, uint256 userId, uint256 machineId, bytes32 sharedPassword) public pure
	    returns (bytes32)
	{
		return keccak256(abi.encode(userName, userId, machineId, sharedPassword));
	}

	function generateOTP(string memory userName, uint256 userId, uint256 machineId, bytes32 sharedPassword) public view
		returns (bytes32 , uint)
	{
		bytes32 userHash = generateUserHash(userName, userId, machineId, sharedPassword);
		// Contatenate user + machineId + sharedPassword + currentBlockNumber
		bytes32 hashToSign = keccak256(abi.encode(userHash, block.number));

		return (hashToSign, block.number);
	}
}
