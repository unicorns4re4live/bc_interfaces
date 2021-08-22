pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;
contract Change {

    struct User {
        uint balance;
        bytes32 pwhash;
        bool role;
    }
    
    // получить информацию о пользователе в системе по его логину - eth-адресу
    mapping(address => User) public users;
    // получить список пользователей в системе
    address[] public userlist;
    
    struct BoostOffer {
        address to_boost;
        address[] proc;
        address cons;
        bool finished;
    }
    
    // получить информацию по голосованию за нового админа, доступен по индексу - id голосования, голоса ЗА не отображаются, для этого функция get_proc()
    mapping(uint => BoostOffer) public boosts;
    // количество голосований за избрание нового админа
    uint public offer_amount;
    
    // получить количество администраторов в системе
    uint public admin_amount;
    address[] zeroArray;
    
    struct Transfer {
        address adr_from;
        address adr_to;
        uint value;
        uint category;
        string description;
        bytes32 pwhash;
        uint time;
        bool finished;
    }
    
    // получить информацию о переводе, доступен по индексу - id перевода
    mapping(uint => Transfer) public transfers;
    // получить количествыо переводов в системе
    uint public transfer_amount;
    
    // наименования категорий, доступны по индексу в массиве
    string[] public categories;
    
    struct Template {
        uint category;
        uint value;
    }
    
    mapping(string => Template) templates;
    // наименования шаблонов, доступны по индексу в массиве
    string[] public template_names;
    
    
    constructor() public {
        bytes32 pwhash = keccak256("123");
        // users
        users[0x5B38Da6a701c568545dCfcB03FcB875f56beddC4] = User(1000, pwhash, false);
        userlist.push(0x5B38Da6a701c568545dCfcB03FcB875f56beddC4);
        users[0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2] = User(1000, pwhash, false);
        userlist.push(0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2);
        users[0xCA35b7d915458EF540aDe6068dFe2F44E8fa733c] = User(1000, pwhash, false);
        userlist.push(0xCA35b7d915458EF540aDe6068dFe2F44E8fa733c);
        users[0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db] = User(1000, pwhash, false);
        userlist.push(0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db);
        // admins
        users[0x78731D3Ca6b7E34aC0F824c42a7cC18A495cabaB] = User(1000, pwhash, true);
        userlist.push(0x78731D3Ca6b7E34aC0F824c42a7cC18A495cabaB);
        users[0x617F2E2fD72FD9D5503197092aC168c91465E7f2] = User(1000, pwhash, true);
        userlist.push(0x617F2E2fD72FD9D5503197092aC168c91465E7f2);
        
        admin_amount = 2;
        
        categories.push("0 - Личный перевод");
        categories.push("1 - Оплата аренды жилья");
        categories.push("2 - Личные взаиморасчеты");
        
        templates["Подарок 10"] = Template(0, 10);
        templates["Подарок 30"] = Template(0, 30);
        templates["Подарок 50"] = Template(0, 50);
        templates["Квартплата 70"] = Template(1, 70);
        templates["Квартплата 90"] = Template(1, 90);
        templates["Погашение задолженности 100"] = Template(2, 100);
        
        template_names = ["Подарок 10", "Подарок 30", "Подарок 50", "Квартплата 70", "Квартплата 90", "Погашение задолженности 100"];
    }
    
    modifier onlyAdmin() {
        require(users[msg.sender].role, "error: you are not admin");
        _;
    }
    
    // User management 
    
    // получить список пользователей
    function get_userlist() public view returns(address[] memory) {
        return(userlist);
    }
    
    // зарегистрировать пользователя, параметры - хэш пароля, 32 байта, алг. - keccak256
    function create_user(bytes32 pwhash) public {
        require(users[msg.sender].pwhash == 0, "error: already registered");
        users[msg.sender] = User(0, pwhash, false);
        userlist.push(msg.sender);
    }
    
    // получить количество пользователей в системе
    function get_users_amount() public view returns(uint) {
        return userlist.length;
    }
    
    // Admin management
    
    // предложить пользователя на роль админа, сразу засчитывается первый голос ЗА, параметр - адрес предлагаемого пользователя
    function offer_user_to_boost(address user_adr) public onlyAdmin {
        require(!users[user_adr].role, "error: already admin");
        boosts[offer_amount] = BoostOffer(user_adr, zeroArray, address(0), false);
        boosts[offer_amount].proc.push(msg.sender);
        offer_amount += 1;
    }
    
    // служебная
    function check_offer(uint offer_id) public onlyAdmin {
        /*
        require(!boosts[offer_id].finished, "error: already finished");
        */
        if (boosts[offer_id].proc.length == admin_amount) {
            boosts[offer_id].finished = true;
            address user_adr = boosts[offer_id].to_boost;
            users[user_adr].role = true;
            admin_amount += 1;
        }
    }
    
    modifier checkVote(uint offer_id) {
        bool vote = false;
        for (uint i = 0; i < boosts[offer_id].proc.length; i++) {
            if (boosts[offer_id].proc[i] == msg.sender) {
                vote = true;
                break;
            }
        }
        /*
        if boosts[offer_amount].cons == msg.sender {
            vote = true;
            break;
        }
        */
        require(!vote, "error: you've been already vote");
        _;
    }
    
    // голосовать ПРОТИВ, аргумент - id голосования, если кто-то голосует против, то голосование сразу завершается, т.к. нужны все голоса ЗА
    function vote_against(uint offer_id) public onlyAdmin checkVote(offer_id) {
        require(!boosts[offer_id].finished, "error: already finished");
        boosts[offer_id].cons = msg.sender;
        boosts[offer_id].finished = true;
    }
    
    // голосовать ЗА, аргумент - id голосования, если это был последний голос ЗА, то голосование заканчивается, пользователь становится Администратором
    function vote_for(uint offer_id) public onlyAdmin checkVote(offer_id) {
        require(!boosts[offer_id].finished, "error: already finished");
        boosts[offer_id].proc.push(msg.sender);
        check_offer(offer_id);
    }
    
    // получить список тех, кто проголосовал ЗА, аргумент - id голосования
    function get_proc(uint offer_id) public view returns(address[] memory) {
        return(boosts[offer_id].proc);
    }
    
    // Category management
    
    // Создать новую категорию, аргумент - имя категории, формат "0 - Имя"
    function create_category(string memory name) public onlyAdmin {
        categories.push(name);
    }
    
    // получить список наименований категорий
    function get_categories() public view returns(string[] memory) {
        return categories;
    }
    
    // Template management
     
    // Создать шаблон, параметры - название нового шаблона, категория перевода, размер перевода
    function create_template(string memory name, uint category, uint value) public onlyAdmin {
        require(category>=0 && category<categories.length, "error: wrong category");
        require(value > 0, "error: wrong value");
        templates[name] = Template(category, value); 
        template_names.push(name);
    }
    
    // получить информацию о шаблоне, параметр - имя шаблона, возвращает номер категории и размер перевода
    function get_template_by_name(string memory name) public view returns(uint, uint) {
        return(templates[name].category, templates[name].value);
    }
    
    // получить список наименований шаблонов
    function get_template_names() public view returns(string[] memory) {
        return(template_names);
    }
    
    
    // Transfers 
    
    // создать перевод, аргумент - адрес кому, сумма перевода, категория перевода, описание, хэш пароля, 32 байта, алг. - keccak256 
    function create_transfer(address adr_to, uint value, uint category, string memory description, bytes32 pwhash) public {
        require(users[msg.sender].balance >= value, "error: not enought money");
        require(category>=0 && category<categories.length, "error: wrong category");
        transfers[transfer_amount] = Transfer(msg.sender, adr_to, value, category, description, pwhash, 0, false);
        transfer_amount += 1;
        users[msg.sender].balance -= value;
    }
    
    // использовать шаблон для перевода, аргумент - адрес кому, наименование шаблона, описание, хэш пароля, 32 байта, алг. - keccak256
    function use_template(address adr_to, string memory template, string memory description, bytes32 pwhash) public {
        uint value = templates[template].value;
        require(value != 0, "error: unknow template");
        uint category = templates[template].category;
        create_transfer(adr_to, value, category, description, pwhash);
    }
    
    // отменить свой перевод, аргумент - id перевода
    function stop_my_transfer(uint tr_id) public {
        require(transfers[tr_id].adr_from == msg.sender, "error: not your transfer");
        require(!transfers[tr_id].finished, "error: already finished");
        transfers[tr_id].finished = true;
        users[msg.sender].balance += transfers[tr_id].value;
    }
    
    // получить перевод, аргумент - id перевода, пароль
    function get_my_transfer(uint tr_id, string memory pw) public {
        require(transfers[tr_id].adr_to == msg.sender, "error: transfer is not for you");
        require(!transfers[tr_id].finished, "error: already finished");
        bytes32  pwhash = keccak256(abi.encodePacked(pw));
        require(pwhash == transfers[tr_id].pwhash, "error: wrong password");
        transfers[tr_id].finished = true;
        transfers[tr_id].time = block.timestamp;
        users[msg.sender].balance += transfers[tr_id].value;
    }
    
}