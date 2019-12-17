pragma solidity >=0.4.22 <0.7.0;

contract Pamette {
/*
    
    function setUser() void public {
    }
    }*/
    
    function isAuthorized(bytes data, bytes signature) public view
        returns (bool)
    {
        return true;
    }
    
    function generateOTP() public view
        returns (string)
    {
        return "password";
    }
}
