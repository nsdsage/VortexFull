// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;
import "../libraries/math/safemath.sol";

/**
 * The ENS registry contract.
 */
contract bdnsRegistry {
    using SafeMath for uint256;

    uint256 _tokenIndex = 0;
    // Mapping from token ID to owner
    mapping (uint256 => address) internal tokenOwner;

    // Mapping from owner to Token ID
    mapping (address => uint256) internal ownedTokens;

    // Mapping from owner to number of owned token
    mapping (address => uint256) internal ownedTokensCount;

    mapping (uint256 => string) internal domainOwnership;

    struct Domain {
      string ip_address;
      string device_type;
   }

    // Map domain to IP Address
    // mapping (string =>  string) internal domains;

    // Map domain to Multiple IP Addresses
    mapping (string => mapping(int256 => Domain[])) internal domains;

    function createRecord(string calldata domain, string calldata ip_address, int256 seclevel, string calldata device_type) external virtual {
        _tokenIndex = SafeMath.add(_tokenIndex, 1);
        domainOwnership[_tokenIndex] = domain;
        domains[domain][seclevel].push(Domain(ip_address, device_type));
        _mint(msg.sender, _tokenIndex);
    }

    function _mint(address _to, uint256 _tokenId) virtual internal {
        require(_to != address(0));
        addTokenTo(_to, _tokenId);
        emit Transfer(address(0), _to, _tokenId);
    }
    function getVersion()public pure returns(uint256) {
        return 1;
    }
    function addTokenTo(address _to, uint256 _tokenId) virtual internal {
        require(tokenOwner[_tokenId] == address(0));
        tokenOwner[_tokenId] = _to;
        ownedTokens[_to] = _tokenId;
        ownedTokensCount[_to] = ownedTokensCount[_to].add(1);
    }
    function getRecord(string memory domain, int256 seclevel) public view returns (Domain[] memory) {
        return domains[domain][seclevel];
    }
    function getOwner(address token) public view returns (uint256) {
        return ownedTokens[token];
    }
    event Transfer(
        address indexed _from,
        address indexed _to,
        uint256 indexed _tokenId
    );
}