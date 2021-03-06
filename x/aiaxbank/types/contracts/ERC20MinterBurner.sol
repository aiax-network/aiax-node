// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0 <0.9.0;
import '@openzeppelin/contracts/access/Ownable.sol';
import '@openzeppelin/contracts/token/ERC20/ERC20.sol';

/**
 * @dev {ERC20} token, including:
 *
 * The account that deploys the contract can mint and burn.
 */
contract ERC20MinterBurner is Ownable, ERC20 {
    uint8 private _decimals;

    /**
     * @dev
     *
     * See {ERC20-constructor}.
     */
    constructor(
        string memory name,
        string memory symbol,
        uint8 decimals_
    ) ERC20(name, symbol) {
        _decimals = decimals_;
    }

    function decimals() public view virtual override returns (uint8) {
        return _decimals;
    }

    /**
     * @dev Creates `amount` new tokens for `to`.
     *
     * See {ERC20-_mint}.
     *
     * Requirements:
     *
     * - the caller must be the account that deployed the contract
     */
    function mint(address to, uint256 amount) public virtual onlyOwner {
        _mint(to, amount);
    }

    /**
     * @dev Burns `amount` tokens from `from`
     *
     * See {ERC20Burnable-_burn}.
     *
     * Requirements:
     *
     * - the caller must be the account that deployed the contract
     */
    function burn(address from, uint256 amount) public virtual onlyOwner {
        _burn(from, amount);
    }
}