package eth

/*
Available Accounts
==================
(0) 0x52C476751142Ce2FB4DB4f19B500e78FEEe10B06 (100 ETH) // OWNER OF NODELIST
(1) 0xFF364b6b86eA5a4f59Cc4989dA23b833daC15304 (100 ETH) FIRST 5 NODES
(2) 0xDC0dd04AaC998E8aA9F2de236b3bA04DDaFd26Ca (100 ETH)
(3) 0x253db77F1aE216722b2F67f33Ef3c8e00B2689e6 (100 ETH)
(4) 0x271346169993368F94cB2c443B8b8CdbDD5EdF04 (100 ETH)
(5) 0xa0ae28Ec27FEa7A577B21330f6CE8aE45A55Fe76 (100 ETH)
(6) 0xf34a875cfFe643D44546b76F0c9412dFb9d2b379 (100 ETH) SECOND 9 NODES
(7) 0x35d946c9c4598CD2eAEe5754CE2041911dc816Ce (100 ETH)
(8) 0xd6ee5E06Ac11a62fd0bE1912dEBEEB4abc24F723 (100 ETH)
(9) 0x40fA4B9e4411E7F5f58713efF426CAD4F0294Ab5 (100 ETH)
(10) 0x0cDa757357158e4d8ad94433e36f1Fe05A1dC576 (100 ETH)
(11) 0xa22e3c16264dc688107142776139d1fB4BB9d549 (100 ETH)
(12) 0x0b998B7229BFd254Acf50b4e2739E73D937Dc1c9 (100 ETH)
(13) 0xFC54C26E24b4570590c11486bD627Aa4B7339523 (100 ETH)
(14) 0xb572081928B988ABE713ffe60F8cf28ef80Eee07 (100 ETH)
(15) 0xd54E0C310a97916E67d07aa501F74524e82c3af1 (100 ETH) THIRD 5 NODES
(16) 0xABA31E255B490365584a56F4EbC5037963E584D5 (100 ETH)
(17) 0x3EcEFaFEa7dB9d0E26dc0D266504587cb66f6008 (100 ETH)
(18) 0x184b56D50300b4cd604A587491cb7BcB0Ffc7454 (100 ETH)
(19) 0xD6eca392adA22e18c9EEbdE2828b38E66813af5f (100 ETH)
(20) 0x3e5d0cd54109F869e5c81A3201F1690aF8A93BBB (100 ETH) THIS IS SIGNER
(21) 0xE90CE283bAFcb6bE825963649c6909ed12F4F2b9 (100 ETH)
(22) 0x1D2daC0166Ea6bBEabf6D36b14782965c165F7DA (100 ETH)
(23) 0x7ca8D75Ec9D7Bed44D619F571c2c82c2687c260D (100 ETH)
(24) 0xa069784Cb3F3666850Fd164E9ECccb05552225eB (100 ETH)
(25) 0xF49091841a56938fF64636d4DF5a54f9abe791B9 (100 ETH)
(26) 0xBce291641cAfD519b222F9860536AFF22D7f41de (100 ETH)
(27) 0x7a438953a3F13827093781aD5D93C1334bf3286C (100 ETH)
(28) 0x55a349e31C0796D2342C99553A1Df664172aDA16 (100 ETH)
(29) 0xE7a7dD62e1060D5BC2CaA08879f8FfFf26ad531E (100 ETH)
(30) 0x7B8D7aC10eD1837E8eA989E94A367cd5Bcdc67AB (100 ETH)
(31) 0x55708b969d1b932f3eE4D99Ff2710bC149bb4A4c (100 ETH)
(32) 0x152cbeB4B29b61FA4C9cBbF90949AaE2f633d357 (100 ETH)
(33) 0x4edA1A0957648532629a5bC132a4e4775326D891 (100 ETH)
(34) 0x3CFd4ECBf805654A32dE0FD243FccAc13B84F054 (100 ETH)
(35) 0x914b92661dF04f50bb8c77e9e8e8BeCBbb7DbA71 (100 ETH)
(36) 0x7Adb9e13Aedcfd787065a82FBf157dA418F926de (100 ETH)
(37) 0xa0273A29dA45854384B95c38cb3Cdeb0F518aa32 (100 ETH)
(38) 0x2DCC25Df05Ac59e7C37B27454442E0026Ddfe06c (100 ETH)
(39) 0x2b78101c096aC2EF1c5539e5e024C4B468415fBD (100 ETH)
(40) 0x157d703420D2dad3CaCb4788fFc2e3b6011050c8 (100 ETH)
(41) 0x6831106504Fa5881DB7CE020fa16CCC9C24A42C5 (100 ETH)
(42) 0x6BC9A29E343f5AEA2F5063c317B967E1d111e41E (100 ETH)
(43) 0xd5bE2B8De1c75137829F9b43b291C69A4858f44e (100 ETH)
(44) 0x5fdEA3d3Edb63E053DD00180686B429A3241F930 (100 ETH)
(45) 0x9831347513C6404DdB3Fad1bCbA20f1Dd17960BF (100 ETH)
(46) 0xd3f772B96C942aE92AC8d523be94D09211E9fC51 (100 ETH)
(47) 0xe61DE55290698aace5DD7EF5B85B8aa890546c37 (100 ETH)
(48) 0xd1Bb0C51D88951019efe948722E852C68d764D66 (100 ETH)
(49) 0xa362781D1B3D15b926F4B64282803D5163a9ADED (100 ETH)
(50) 0x0D8aE13a17E6A411e56752b6c339B96c38b492f6 (100 ETH)
(51) 0x396CbB9C2C88E4749DcaB279b04e55fB21110766 (100 ETH)
(52) 0x0684e5d8766b12FE0248e314230A6b9560eedad3 (100 ETH)
(53) 0x67a4Ee68345BaA9cab85DEe8aB980f57d6D8bDfd (100 ETH)
(54) 0xFeFC970EC697962DB8dD2326e26940f657Ee791B (100 ETH)
(55) 0x19EbeC1014DCB459d9a9422A7178811c9a083fB2 (100 ETH)
(56) 0xCd89161BD7b94858431178Ba168E291b996B991E (100 ETH)
(57) 0xE5Aa21B7307c48fE8Ec59530D1Ff2aec0AD42A2c (100 ETH)
(58) 0x1F87368FB479748C5d270633F9493614c9903cEc (100 ETH)
(59) 0x29627Bd7e9283c0FF670188e2E60658522d742E8 (100 ETH)
(60) 0xEDEdC4887F7c590B59b74A4dDe03d34B405Ff54C (100 ETH)
(61) 0x23464B33A8463842907C5A71997Ffb71448DDE48 (100 ETH)
(62) 0xe629Caa91A93fd345afccAf62b1b93E8DFe4567F (100 ETH)
(63) 0x3434b90759f877A8005589447cd0bCD4337C4351 (100 ETH)
(64) 0x5e8229027B2F7f9D249470355DFfbe225109BB16 (100 ETH)
(65) 0x8C4737915ad701e04C4A68Fb58875289d33404f1 (100 ETH)
(66) 0x3FeB5DAf29d3193b3ceba24015980D4941370F0E (100 ETH)
(67) 0x9AF4665009002F08C3267705B979ba9eA72f5444 (100 ETH)
(68) 0xB7BDD0c7896aC6E9D8C0A22840da5D2Ffda49A68 (100 ETH)
(69) 0x091Fcb985dC43A50FED245d19a0B50569F0Ba738 (100 ETH)
(70) 0xC4B6b93f0979631296Ef7f92247221943f0382d1 (100 ETH)
(71) 0xf1EBF4fe999f87d13525F6a5474E289f4ec3660E (100 ETH)
(72) 0xA9Cf0e8c99D1Fdf7497BcDde69f1B72A45D1611f (100 ETH)
(73) 0x7b3a1C225e32E1725654AF18D2bca0A34a1da2ED (100 ETH)
(74) 0x5d5BBD253a1De95c1D7D8D2aC770147fA67864f0 (100 ETH)
(75) 0xba19183208738f3BC23f44F28f7c2f4aC408db83 (100 ETH)
(76) 0xd2777017e15CbEda38e6A3f136AE62B0fa8a8ba1 (100 ETH)
(77) 0xfb64F3BD2BEf3e0491ac61b3b77cD9c4F785079D (100 ETH)
(78) 0x8b359E004c8071B6940BDeF83cd87401ecEAa14c (100 ETH)
(79) 0x0a33bfa0BbA362404Da3aeA744cc8275035d849B (100 ETH)
(80) 0x0E2E2773b5d7ec57E4785C757F495a25634768A6 (100 ETH)
(81) 0x9d7D0C9D3D2Df28751026552B877A879964f79C9 (100 ETH)
(82) 0x059AF2501EC4Fab7a11A85f5C91FDc1a60B1A487 (100 ETH)
(83) 0x77A2C3F2c4c35b4B4aFA7235c888eeB54DbECe26 (100 ETH)
(84) 0x343C206d64569b500fF756a354567c5CB648686d (100 ETH)
(85) 0x99B4764C28C9188B40Aad745B80320175c421a89 (100 ETH)
(86) 0xb5BFD363c78cD899F4fdecDB4DB98e2032fE8A65 (100 ETH)
(87) 0xA1630bA9b7Ca842226fEFfb3A61E9cD414E0D326 (100 ETH)
(88) 0x0a7b3D284D39F8a2f8166f0D95f115a9799b23D7 (100 ETH)
(89) 0x208C52807B9D946F897D6c8D3796ABf93af56Dc2 (100 ETH)
(90) 0x348ad3696F44225039746e3272b5Ab7c5c87F71C (100 ETH)
(91) 0xF5BD4dD32cFd0d96f5c6fC3037BB7C50D52B5022 (100 ETH)
(92) 0xeE134841575a3CEceAA19f181CaFeA2B70606ce0 (100 ETH)
(93) 0x9123fB497ABb5ee6981C3154d5ca5E52B6F073C1 (100 ETH)
(94) 0xcd244A9E91197278935368c49480cD9fb5091e99 (100 ETH)
(95) 0x8aF0d48849F9D416a13a79Ed936Dc8Af2F2C899a (100 ETH)
(96) 0x9203986a453d809EDF6ee18989f7Bf20f1F92d57 (100 ETH)
(97) 0xFAB089AFA322310C85d82f56e504eC545a6b71FE (100 ETH)
(98) 0x955bcd0762cA61f62cB04BE89BAe2a66d24e0f54 (100 ETH)
(99) 0xCF11be3bc60E44AcE644F92B81B50d9fcf4e73dd (100 ETH)

Private Keys
==================
(0) 0x2cc22a5c9b07f1286d7e708ea326b4a0ac18a70decfc320176cf5f103a5732ca
(1) 0x29909a750dc6abc3e3c83de9c6da9d6faf9fde4eebb61fa21221415557de5a0b
(2) 0x318437152af791c195d32218ea759e10b8ac5ac78e1cb2e131997a2d682c6613
(3) 0x74e31abf0d78420fb095f6e82ad8a0d2158c92079807c46661ea463df3ad53b5
(4) 0x798c8d2e1f1f3de5ffa27350f481384385be06e21833c6c5988279e86fdf93c4
(5) 0x83781f4a87b3a1fd7473652afc441f4abb111a46c44b404ec640f43b89202387
(6) 0x1ef08a4b9fd575911d5f7b67eaa81e5042012bfc3c9cd374733b99febff292c8
(7) 0x39466419d15e0c79ca53f8e20fff9c8da938674f4e7446d7e673e5bb3093219e
(8) 0xc605a67c38a0e6169e8d86cc4b4fba6fea217a25193b094094e8625a1f4b69f8
(9) 0x2623df0903c45d08e7d557dd94ce7fdaca5134beb6c7694de27e7e6de111ada2
(10) 0xbee504c30fa61346211df674f0011450b21c37e8045c11aa0fe2609a0a75ad40
(11) 0x10677a6dbceebea8ed7803ce2bd4a89703b28582ced65b76e0c8db399ac89dfe
(12) 0xd81e5ddd2c6b900897796e53c788384eb0a243dc385079116f6ddb4cba81b342
(13) 0x45d719d07ec2b9e478b57b4b5ecf2382ee778644084935c7e528e16333d15b77
(14) 0x8e0fbabca4728106a591256628c651dd1db173b7341363e94092c648d9a7539b
(15) 0x8dcd8d195945b2b77689d1c7a5fcaba03bfcc146e114640036a0ad73e0c92ccc
(16) 0x1e69eed579adf1b53b0881beaa131ced8f132c1e01aabc06be7c8a942ee07f36
(17) 0x8205074187301bd7c12c4d8b90a7984041ddfcef664172403eee983dcd61c816
(18) 0xd1e99f1e00bb0e60caa7b92a724d5532910a98f3da5e2a4953a9e790d9ec018e
(19) 0x59fcac109ac412fbda8acfb68fd1d9a0b1b22d53c9b15802a966d1cd6f367b57
(20) 0x15ca0f59d9526f15bb02cad680be2d5eb3de8a50069e8e2671aaa1a583e0aca5
(21) 0x078fede9e87c96e7af42d54c8eb4599269b67d1d376333a3a86ab21b2dc4fcb1
(22) 0xa429eade3c64c43e48ce6b3893a68ff6061279817e202f167f51f22cdb3fa679
(23) 0xe94f68c0ef6574602966dd91160cea6f8c8c77be8c97cda6601e80580dd1f135
(24) 0x52a89b45fd72079fd138ac83a82e7e06d37303912d28b51d912e96df415b5ff3
(25) 0xf5d6c3e64c09311197033dfff2a5d2b51086d8ef042742504bee9470f31d6e33
(26) 0x5e03ddc2c8ec8417a8513f19c298c82716296a80da505700e5c2d5dd05810fdd
(27) 0x152ecb01c702d04d52ed06972cf770681c3135f31b684728328e4380678d8afe
(28) 0x09642f2068001c173b53aa731fb41a6600ccb5d3989cb03a2ff0351eef1660b2
(29) 0x162cc66411e9f82d6234a099144f3f10375344d7f56841bcb4e7dce5aec8f16f
(30) 0x2f05f01fc47be70a76b3fc18d81a27c64972515357e40707cbe6a545e0ad0f02
(31) 0x13001d5682452887b7529527d0b67d26be116a56f289fbd1c5f93a101a5da608
(32) 0x7d128ae93218892c89c3c2b6959e4f9c905e4107c5357ed52f338c7c07cbf434
(33) 0x9059d899cc188378a4b8d2af09375d0b099fb59a6ee3e9d3d3dc56758de397b1
(34) 0xed6982d0a404241cb1e0a1883c481a8f81fea78c2b1aac53feb6b6163fe78637
(35) 0xa2d65045053c225b86f757ba405c7b28b86f16f30634f4b298819b495abac921
(36) 0xffbcc45427b2609e06f973db53d5e37687b98a3b6f8d8842c4f1b7e221f9cce3
(37) 0x5751934efc7028d0710dbdcf41ef5681961a5d7ce5ad6abf492190c0233bc13d
(38) 0xe76e5f5b37f9078cf6294180c82d97d650c949cca84a57a428294f818e8f50d7
(39) 0x76d17879259989e763de06c1887d9fe30e4b24271f703934808c76ba8cc2b2a7
(40) 0xdc359090d20de81d14d3459695b9addf5bcb847d4cb84355146598e6a83a380d
(41) 0x9d77ab4c62e518414f709f9669f7a5f7925a0f673afec0c45cbf8564a0837499
(42) 0x79ee01b4b0b14f930120dbf35c968bf44c3d4ccf06961ebd337fbf32f11e6c3b
(43) 0x62396b984999a029c409d11b101b6a86c06c17ff7c58b8ed54117dc5105c5cf4
(44) 0x86512b06d823d4033c3b3c8e55f44c1745f1a85e0cf966087f2f33e03cc43b14
(45) 0x17663afc26b3e39826a680f395c462fd44ee073754741ced9d6aab9a4edc48f3
(46) 0x748d455cddf7387ab07868ee7f6020c2fd0aaa0b4cc102f9fc912a3cffffd5c1
(47) 0x3973a3505cc573354b0cef2c12a3a00129d936851ecee33dbafed49fea1ff5c9
(48) 0xe4746ee3eb6efb80e4e8dd66df626bf38be2d86ea3c1e51a6bb4a879e47c6804
(49) 0x1eacce98d8b9c393927a8f0367b6364f94b7ae6ace254ccaaeab07cb4386c051
(50) 0xb18208b0ea8931d22fbe21db95fe5ee4ef4bb1b3899760943d172cb453f02851
(51) 0xca2d7d41f10cbc6604e07d1b7cb02af0f8a21ccb7eb3c08edfe8025e6ea68197
(52) 0x2b58121b2ef28e6ffd8e43b272c49ac43d8a51034148fc46d1004a16c88af2c6
(53) 0x19a97f583e28ad2ed66536cc3e35e682ff681be41965cae4c7f8d2dbdbfd726a
(54) 0x18fd10104c36716f9175a9de0e3774826e8ae4e8ac76dfcb65ee1e0f1e66bc47
(55) 0x88ed0fbcdf07f7ca3dd32ec00d3e6098c4400c546a5cc3c517bb0efc2242cf79
(56) 0x143380727a1fad6de0c11f97c81fde8abc033f2603be8740dd93d509e726e217
(57) 0x84a30df8aeb8ad08474e1792c4bb7752a901fa98e66856588902417e7ed63063
(58) 0x329764f80bedf0f343711c9ea756c3b7cc635e2cee72b8f5d7283f09219b3f30
(59) 0x2a9b73a3d9dee56d35316ce4a79bee9e9fc600b27cf7230775f7e740236d6ce0
(60) 0xe363ecdf669de1c5e6e308c5d845a1b7c5e984f7fc9e6e6f161cbacdcff3f3fc
(61) 0xd07e24ffa9920dcc4fadf7083df816ee65256636271a98ab7a5898dbae11bc11
(62) 0x60eac59b78cbd8315b5d18c3807aa077928483c4399ec50e34c447cf3d8e48c7
(63) 0x319163b763d0541253f294ce9840bca58599d9259ae2a8324875c41a4eea8c86
(64) 0x8a3647256bdf91020db20c37ba093b07f42bc3a3df2178c4672ee9f7d36653b9
(65) 0xf14c633c1516b3fff8055b92990895e5c2ad49ae61c5b8e2a635c57cf8d9a3dd
(66) 0x3a01d39f0fc9c5aeb9f0911c88c5b0af0355edaa7bc6aef3284ecae40efbf168
(67) 0x59ece72a8276615ae68ad0ef12006ed3d574afd0b14c73a7f0d12027bbbf9cd8
(68) 0x4d788605e7cf932674f92e229173124d7e2a7c529be66f43c69c3c7cb856f1ae
(69) 0x0a59897dc2188aa411c1f5bd8ee6e588b4bc5eaff8e015e0ad39e5a910711285
(70) 0x107a5301982d2915ad659ab12cc76de88ba78dcb60d5ff4dc3c3202aa1b232be
(71) 0x7a43fdf3b689f27b814178ed726e0fea0aa116c0b3213b726545f14ff3d11c01
(72) 0xe1de49153c45ad16d57cbcc0c44e6b27fa765710f599a019b77d75e72b3aba8d
(73) 0x5217ccf2c6355e388858dbaba3cfc6d5d95f425b352f1cb0d0ed045fa7ad1cdf
(74) 0x02eb3bd1cb4ab82aa6bc04025f7195bea8a3a7dd42b7d4742fdeb61cc1b35cab
(75) 0x41abf3506521c2a33522cc85440504ceeaf957edfa38f2f5a43a6b2ecf9241f0
(76) 0x1293899f1270ffff9a329cf0703e2107f482a4d94c7a7a7b2bf1d8853516b36b
(77) 0xeaed486c84e6a12179a73bb51b0e023b2a5807d05587ce8414a75eae726780d3
(78) 0x764f3fcb77c4c06ababbb7bc5e6a840569119eae20fb6e235c2803557d7f8e28
(79) 0x04dc5b83005ea1cac04f7d8fb8b33b5437f9a332e4a8a624ad823962ba4c12b5
(80) 0x6324ccea71418a28088edcb56e1ab145b7ff46e37770120377f6c8da14b5dfe7
(81) 0xdb0813b77413dd4832cc5a268d107b13f8cc0e4f543fea0870093486169efe8f
(82) 0x30ef44eea8b84b714dc5e2ba811c2ca9564fc88561dca693942b867f139ebe88
(83) 0x5ca4fe6939239b7d965b40431c393dda5eb4c76de78da65db8a476260c80a9e9
(84) 0xe87f8aa5ac08ef91d87684a0eb27745274096da415ac373d3118d3eeb248cb79
(85) 0x7224cbdd8f3652272ab91fcb52ab698d6e17458ca7ddf41597a0f47eb0d479a8
(86) 0x7f8dd45196bae5aa266c23c1c2136d722b5094f3b2f048f1fff1df9f689e4f82
(87) 0x67cf4ca479b33ac82abcd5a6e5c9cac0e0ee6bc7f7609a2a1282b5c9b56c13da
(88) 0x022dbaafdb118c47159691be6bae332c4b1f68b2ad1f7ac4ed0d8437b44ed4a2
(89) 0x9f73af306d197919d8120dbecffcbbecdf33d2a37baec17042a551ff41985507
(90) 0x340aa63e398e03f1cd39c97502ca783dc32786c39f05dafdef54155c38847542
(91) 0x69b480d3a41488c56105f1e41141648f79fab067ab1379c36efb6686260a3da9
(92) 0x150349a450a84d3ca36de71dc6c5fdd025ef93901ddab5eeb177e318405486b1
(93) 0xb5a3e294dbc4b94556ff2cad7a75e6da2e53ac41244e305ae2512bf4d6b4f511
(94) 0x598c06f72fc6c5bdb558b1a59fc53fe8143043dfaf7f708f5e220c46aad3cab5
(95) 0x2bf02ea6e3e1f50b3853294183bcd51566ea1ef741e94d064abf7693424ec739
(96) 0xfbe403773744ab020f995a239889877cc1f438c6861e935b2d2fd5c760a9088a
(97) 0x25281afef03a84fe4902a59132fe305051674f4715feb9bb8fb6609c055e7223
(98) 0x4bc8abb1a95bbe73c850697cb7bd2fc59c8eaa50336a4334484c95a931c55038
(99) 0xc348c95ff57c40fe57bfd97b671b7a5f90421f83ee321d984f6ddf89c45262ef
*/

// GanacheKeys - log of all the keys
type GanacheKeys struct {
	PublicAddress map[int]string
	PrivateKeys   map[int]string
}

var TestKeys = GanacheKeys{
	PublicAddress: map[int]string{
		0:  "0x52C476751142Ce2FB4DB4f19B500e78FEEe10B06", // OWNER OF NODELIST
		1:  "0xFF364b6b86eA5a4f59Cc4989dA23b833daC15304", // FIRST 5 NODES
		2:  "0xDC0dd04AaC998E8aA9F2de236b3bA04DDaFd26Ca",
		3:  "0x253db77F1aE216722b2F67f33Ef3c8e00B2689e6",
		4:  "0x271346169993368F94cB2c443B8b8CdbDD5EdF04",
		5:  "0xa0ae28Ec27FEa7A577B21330f6CE8aE45A55Fe76",
		6:  "0xf34a875cfFe643D44546b76F0c9412dFb9d2b379", // SECOND 9 NODES
		7:  "0x35d946c9c4598CD2eAEe5754CE2041911dc816Ce",
		8:  "0xd6ee5E06Ac11a62fd0bE1912dEBEEB4abc24F723",
		9:  "0x40fA4B9e4411E7F5f58713efF426CAD4F0294Ab5",
		10: "0x0cDa757357158e4d8ad94433e36f1Fe05A1dC576",
		11: "0xa22e3c16264dc688107142776139d1fB4BB9d549",
		12: "0x0b998B7229BFd254Acf50b4e2739E73D937Dc1c9",
		13: "0xFC54C26E24b4570590c11486bD627Aa4B7339523",
		14: "0xb572081928B988ABE713ffe60F8cf28ef80Eee07",
		15: "0xd54E0C310a97916E67d07aa501F74524e82c3af1", // THIRD 5 NODES
		16: "0xABA31E255B490365584a56F4EbC5037963E584D5",
		17: "0x3EcEFaFEa7dB9d0E26dc0D266504587cb66f6008",
		18: "0x184b56D50300b4cd604A587491cb7BcB0Ffc7454",
		19: "0xD6eca392adA22e18c9EEbdE2828b38E66813af5f",
		20: "0x3e5d0cd54109F869e5c81A3201F1690aF8A93BBB", // THIS IS SIGNER
		21: "0xE90CE283bAFcb6bE825963649c6909ed12F4F2b9",
		22: "0x1D2daC0166Ea6bBEabf6D36b14782965c165F7DA",
		23: "0x7ca8D75Ec9D7Bed44D619F571c2c82c2687c260D",
		24: "0xa069784Cb3F3666850Fd164E9ECccb05552225eB",
		25: "0xF49091841a56938fF64636d4DF5a54f9abe791B9",
		26: "0xBce291641cAfD519b222F9860536AFF22D7f41de",
		27: "0x7a438953a3F13827093781aD5D93C1334bf3286C",
		28: "0x55a349e31C0796D2342C99553A1Df664172aDA16",
		29: "0xE7a7dD62e1060D5BC2CaA08879f8FfFf26ad531E",
		30: "0x7B8D7aC10eD1837E8eA989E94A367cd5Bcdc67AB",
		31: "0x55708b969d1b932f3eE4D99Ff2710bC149bb4A4c",
		32: "0x152cbeB4B29b61FA4C9cBbF90949AaE2f633d357",
		33: "0x4edA1A0957648532629a5bC132a4e4775326D891",
		34: "0x3CFd4ECBf805654A32dE0FD243FccAc13B84F054",
		35: "0x914b92661dF04f50bb8c77e9e8e8BeCBbb7DbA71",
		36: "0x7Adb9e13Aedcfd787065a82FBf157dA418F926de",
		37: "0xa0273A29dA45854384B95c38cb3Cdeb0F518aa32",
		38: "0x2DCC25Df05Ac59e7C37B27454442E0026Ddfe06c",
		39: "0x2b78101c096aC2EF1c5539e5e024C4B468415fBD",
		40: "0x157d703420D2dad3CaCb4788fFc2e3b6011050c8",
		41: "0x6831106504Fa5881DB7CE020fa16CCC9C24A42C5",
		42: "0x6BC9A29E343f5AEA2F5063c317B967E1d111e41E",
		43: "0xd5bE2B8De1c75137829F9b43b291C69A4858f44e",
		44: "0x5fdEA3d3Edb63E053DD00180686B429A3241F930",
		45: "0x9831347513C6404DdB3Fad1bCbA20f1Dd17960BF",
		46: "0xd3f772B96C942aE92AC8d523be94D09211E9fC51",
		47: "0xe61DE55290698aace5DD7EF5B85B8aa890546c37",
		48: "0xd1Bb0C51D88951019efe948722E852C68d764D66",
		49: "0xa362781D1B3D15b926F4B64282803D5163a9ADED",
		50: "0x0D8aE13a17E6A411e56752b6c339B96c38b492f6",
		51: "0x396CbB9C2C88E4749DcaB279b04e55fB21110766",
		52: "0x0684e5d8766b12FE0248e314230A6b9560eedad3",
		53: "0x67a4Ee68345BaA9cab85DEe8aB980f57d6D8bDfd",
		54: "0xFeFC970EC697962DB8dD2326e26940f657Ee791B",
		55: "0x19EbeC1014DCB459d9a9422A7178811c9a083fB2",
		56: "0xCd89161BD7b94858431178Ba168E291b996B991E",
		57: "0xE5Aa21B7307c48fE8Ec59530D1Ff2aec0AD42A2c",
		58: "0x1F87368FB479748C5d270633F9493614c9903cEc",
		59: "0x29627Bd7e9283c0FF670188e2E60658522d742E8",
		60: "0xEDEdC4887F7c590B59b74A4dDe03d34B405Ff54C",
		61: "0x23464B33A8463842907C5A71997Ffb71448DDE48",
		62: "0xe629Caa91A93fd345afccAf62b1b93E8DFe4567F",
		63: "0x3434b90759f877A8005589447cd0bCD4337C4351",
		64: "0x5e8229027B2F7f9D249470355DFfbe225109BB16",
		65: "0x8C4737915ad701e04C4A68Fb58875289d33404f1",
		66: "0x3FeB5DAf29d3193b3ceba24015980D4941370F0E",
		67: "0x9AF4665009002F08C3267705B979ba9eA72f5444",
		68: "0xB7BDD0c7896aC6E9D8C0A22840da5D2Ffda49A68",
		69: "0x091Fcb985dC43A50FED245d19a0B50569F0Ba738",
		70: "0xC4B6b93f0979631296Ef7f92247221943f0382d1",
		71: "0xf1EBF4fe999f87d13525F6a5474E289f4ec3660E",
		72: "0xA9Cf0e8c99D1Fdf7497BcDde69f1B72A45D1611f",
		73: "0x7b3a1C225e32E1725654AF18D2bca0A34a1da2ED",
		74: "0x5d5BBD253a1De95c1D7D8D2aC770147fA67864f0",
		75: "0xba19183208738f3BC23f44F28f7c2f4aC408db83",
		76: "0xd2777017e15CbEda38e6A3f136AE62B0fa8a8ba1",
		77: "0xfb64F3BD2BEf3e0491ac61b3b77cD9c4F785079D",
		78: "0x8b359E004c8071B6940BDeF83cd87401ecEAa14c",
		79: "0x0a33bfa0BbA362404Da3aeA744cc8275035d849B",
		80: "0x0E2E2773b5d7ec57E4785C757F495a25634768A6",
		81: "0x9d7D0C9D3D2Df28751026552B877A879964f79C9",
		82: "0x059AF2501EC4Fab7a11A85f5C91FDc1a60B1A487",
		83: "0x77A2C3F2c4c35b4B4aFA7235c888eeB54DbECe26",
		84: "0x343C206d64569b500fF756a354567c5CB648686d",
		85: "0x99B4764C28C9188B40Aad745B80320175c421a89",
		86: "0xb5BFD363c78cD899F4fdecDB4DB98e2032fE8A65",
		87: "0xA1630bA9b7Ca842226fEFfb3A61E9cD414E0D326",
		88: "0x0a7b3D284D39F8a2f8166f0D95f115a9799b23D7",
		89: "0x208C52807B9D946F897D6c8D3796ABf93af56Dc2",
		90: "0x348ad3696F44225039746e3272b5Ab7c5c87F71C",
		91: "0xF5BD4dD32cFd0d96f5c6fC3037BB7C50D52B5022",
		92: "0xeE134841575a3CEceAA19f181CaFeA2B70606ce0",
		93: "0x9123fB497ABb5ee6981C3154d5ca5E52B6F073C1",
		94: "0xcd244A9E91197278935368c49480cD9fb5091e99",
		95: "0x8aF0d48849F9D416a13a79Ed936Dc8Af2F2C899a",
		96: "0x9203986a453d809EDF6ee18989f7Bf20f1F92d57",
		97: "0xFAB089AFA322310C85d82f56e504eC545a6b71FE",
		98: "0x955bcd0762cA61f62cB04BE89BAe2a66d24e0f54",
		99: "0xCF11be3bc60E44AcE644F92B81B50d9fcf4e73dd",
	},
	PrivateKeys: map[int]string{
		0:  "2cc22a5c9b07f1286d7e708ea326b4a0ac18a70decfc320176cf5f103a5732ca",
		1:  "29909a750dc6abc3e3c83de9c6da9d6faf9fde4eebb61fa21221415557de5a0b",
		2:  "318437152af791c195d32218ea759e10b8ac5ac78e1cb2e131997a2d682c6613",
		3:  "74e31abf0d78420fb095f6e82ad8a0d2158c92079807c46661ea463df3ad53b5",
		4:  "798c8d2e1f1f3de5ffa27350f481384385be06e21833c6c5988279e86fdf93c4",
		5:  "83781f4a87b3a1fd7473652afc441f4abb111a46c44b404ec640f43b89202387",
		6:  "1ef08a4b9fd575911d5f7b67eaa81e5042012bfc3c9cd374733b99febff292c8",
		7:  "39466419d15e0c79ca53f8e20fff9c8da938674f4e7446d7e673e5bb3093219e",
		8:  "c605a67c38a0e6169e8d86cc4b4fba6fea217a25193b094094e8625a1f4b69f8",
		9:  "2623df0903c45d08e7d557dd94ce7fdaca5134beb6c7694de27e7e6de111ada2",
		10: "bee504c30fa61346211df674f0011450b21c37e8045c11aa0fe2609a0a75ad40",
		11: "10677a6dbceebea8ed7803ce2bd4a89703b28582ced65b76e0c8db399ac89dfe",
		12: "d81e5ddd2c6b900897796e53c788384eb0a243dc385079116f6ddb4cba81b342",
		13: "45d719d07ec2b9e478b57b4b5ecf2382ee778644084935c7e528e16333d15b77",
		14: "8e0fbabca4728106a591256628c651dd1db173b7341363e94092c648d9a7539b",
		15: "8dcd8d195945b2b77689d1c7a5fcaba03bfcc146e114640036a0ad73e0c92ccc",
		16: "1e69eed579adf1b53b0881beaa131ced8f132c1e01aabc06be7c8a942ee07f36",
		17: "8205074187301bd7c12c4d8b90a7984041ddfcef664172403eee983dcd61c816",
		18: "d1e99f1e00bb0e60caa7b92a724d5532910a98f3da5e2a4953a9e790d9ec018e",
		19: "59fcac109ac412fbda8acfb68fd1d9a0b1b22d53c9b15802a966d1cd6f367b57",
		20: "15ca0f59d9526f15bb02cad680be2d5eb3de8a50069e8e2671aaa1a583e0aca5",
		21: "078fede9e87c96e7af42d54c8eb4599269b67d1d376333a3a86ab21b2dc4fcb1",
		22: "a429eade3c64c43e48ce6b3893a68ff6061279817e202f167f51f22cdb3fa679",
		23: "e94f68c0ef6574602966dd91160cea6f8c8c77be8c97cda6601e80580dd1f135",
		24: "52a89b45fd72079fd138ac83a82e7e06d37303912d28b51d912e96df415b5ff3",
		25: "f5d6c3e64c09311197033dfff2a5d2b51086d8ef042742504bee9470f31d6e33",
		26: "5e03ddc2c8ec8417a8513f19c298c82716296a80da505700e5c2d5dd05810fdd",
		27: "152ecb01c702d04d52ed06972cf770681c3135f31b684728328e4380678d8afe",
		28: "09642f2068001c173b53aa731fb41a6600ccb5d3989cb03a2ff0351eef1660b2",
		29: "162cc66411e9f82d6234a099144f3f10375344d7f56841bcb4e7dce5aec8f16f",
		30: "2f05f01fc47be70a76b3fc18d81a27c64972515357e40707cbe6a545e0ad0f02",
		31: "13001d5682452887b7529527d0b67d26be116a56f289fbd1c5f93a101a5da608",
		32: "7d128ae93218892c89c3c2b6959e4f9c905e4107c5357ed52f338c7c07cbf434",
		33: "9059d899cc188378a4b8d2af09375d0b099fb59a6ee3e9d3d3dc56758de397b1",
		34: "ed6982d0a404241cb1e0a1883c481a8f81fea78c2b1aac53feb6b6163fe78637",
		35: "a2d65045053c225b86f757ba405c7b28b86f16f30634f4b298819b495abac921",
		36: "ffbcc45427b2609e06f973db53d5e37687b98a3b6f8d8842c4f1b7e221f9cce3",
		37: "5751934efc7028d0710dbdcf41ef5681961a5d7ce5ad6abf492190c0233bc13d",
		38: "e76e5f5b37f9078cf6294180c82d97d650c949cca84a57a428294f818e8f50d7",
		39: "76d17879259989e763de06c1887d9fe30e4b24271f703934808c76ba8cc2b2a7",
		40: "dc359090d20de81d14d3459695b9addf5bcb847d4cb84355146598e6a83a380d",
		41: "9d77ab4c62e518414f709f9669f7a5f7925a0f673afec0c45cbf8564a0837499",
		42: "79ee01b4b0b14f930120dbf35c968bf44c3d4ccf06961ebd337fbf32f11e6c3b",
		43: "62396b984999a029c409d11b101b6a86c06c17ff7c58b8ed54117dc5105c5cf4",
		44: "86512b06d823d4033c3b3c8e55f44c1745f1a85e0cf966087f2f33e03cc43b14",
		45: "17663afc26b3e39826a680f395c462fd44ee073754741ced9d6aab9a4edc48f3",
		46: "748d455cddf7387ab07868ee7f6020c2fd0aaa0b4cc102f9fc912a3cffffd5c1",
		47: "3973a3505cc573354b0cef2c12a3a00129d936851ecee33dbafed49fea1ff5c9",
		48: "e4746ee3eb6efb80e4e8dd66df626bf38be2d86ea3c1e51a6bb4a879e47c6804",
		49: "1eacce98d8b9c393927a8f0367b6364f94b7ae6ace254ccaaeab07cb4386c051",
		50: "b18208b0ea8931d22fbe21db95fe5ee4ef4bb1b3899760943d172cb453f02851",
		51: "ca2d7d41f10cbc6604e07d1b7cb02af0f8a21ccb7eb3c08edfe8025e6ea68197",
		52: "2b58121b2ef28e6ffd8e43b272c49ac43d8a51034148fc46d1004a16c88af2c6",
		53: "19a97f583e28ad2ed66536cc3e35e682ff681be41965cae4c7f8d2dbdbfd726a",
		54: "18fd10104c36716f9175a9de0e3774826e8ae4e8ac76dfcb65ee1e0f1e66bc47",
		55: "88ed0fbcdf07f7ca3dd32ec00d3e6098c4400c546a5cc3c517bb0efc2242cf79",
		56: "143380727a1fad6de0c11f97c81fde8abc033f2603be8740dd93d509e726e217",
		57: "84a30df8aeb8ad08474e1792c4bb7752a901fa98e66856588902417e7ed63063",
		58: "329764f80bedf0f343711c9ea756c3b7cc635e2cee72b8f5d7283f09219b3f30",
		59: "2a9b73a3d9dee56d35316ce4a79bee9e9fc600b27cf7230775f7e740236d6ce0",
		60: "e363ecdf669de1c5e6e308c5d845a1b7c5e984f7fc9e6e6f161cbacdcff3f3fc",
		61: "d07e24ffa9920dcc4fadf7083df816ee65256636271a98ab7a5898dbae11bc11",
		62: "60eac59b78cbd8315b5d18c3807aa077928483c4399ec50e34c447cf3d8e48c7",
		63: "319163b763d0541253f294ce9840bca58599d9259ae2a8324875c41a4eea8c86",
		64: "8a3647256bdf91020db20c37ba093b07f42bc3a3df2178c4672ee9f7d36653b9",
		65: "f14c633c1516b3fff8055b92990895e5c2ad49ae61c5b8e2a635c57cf8d9a3dd",
		66: "3a01d39f0fc9c5aeb9f0911c88c5b0af0355edaa7bc6aef3284ecae40efbf168",
		67: "59ece72a8276615ae68ad0ef12006ed3d574afd0b14c73a7f0d12027bbbf9cd8",
		68: "4d788605e7cf932674f92e229173124d7e2a7c529be66f43c69c3c7cb856f1ae",
		69: "0a59897dc2188aa411c1f5bd8ee6e588b4bc5eaff8e015e0ad39e5a910711285",
		70: "107a5301982d2915ad659ab12cc76de88ba78dcb60d5ff4dc3c3202aa1b232be",
		71: "7a43fdf3b689f27b814178ed726e0fea0aa116c0b3213b726545f14ff3d11c01",
		72: "e1de49153c45ad16d57cbcc0c44e6b27fa765710f599a019b77d75e72b3aba8d",
		73: "5217ccf2c6355e388858dbaba3cfc6d5d95f425b352f1cb0d0ed045fa7ad1cdf",
		74: "02eb3bd1cb4ab82aa6bc04025f7195bea8a3a7dd42b7d4742fdeb61cc1b35cab",
		75: "41abf3506521c2a33522cc85440504ceeaf957edfa38f2f5a43a6b2ecf9241f0",
		76: "1293899f1270ffff9a329cf0703e2107f482a4d94c7a7a7b2bf1d8853516b36b",
		77: "eaed486c84e6a12179a73bb51b0e023b2a5807d05587ce8414a75eae726780d3",
		78: "764f3fcb77c4c06ababbb7bc5e6a840569119eae20fb6e235c2803557d7f8e28",
		79: "04dc5b83005ea1cac04f7d8fb8b33b5437f9a332e4a8a624ad823962ba4c12b5",
		80: "6324ccea71418a28088edcb56e1ab145b7ff46e37770120377f6c8da14b5dfe7",
		81: "db0813b77413dd4832cc5a268d107b13f8cc0e4f543fea0870093486169efe8f",
		82: "30ef44eea8b84b714dc5e2ba811c2ca9564fc88561dca693942b867f139ebe88",
		83: "5ca4fe6939239b7d965b40431c393dda5eb4c76de78da65db8a476260c80a9e9",
		84: "e87f8aa5ac08ef91d87684a0eb27745274096da415ac373d3118d3eeb248cb79",
		85: "7224cbdd8f3652272ab91fcb52ab698d6e17458ca7ddf41597a0f47eb0d479a8",
		86: "7f8dd45196bae5aa266c23c1c2136d722b5094f3b2f048f1fff1df9f689e4f82",
		87: "67cf4ca479b33ac82abcd5a6e5c9cac0e0ee6bc7f7609a2a1282b5c9b56c13da",
		88: "022dbaafdb118c47159691be6bae332c4b1f68b2ad1f7ac4ed0d8437b44ed4a2",
		89: "9f73af306d197919d8120dbecffcbbecdf33d2a37baec17042a551ff41985507",
		90: "340aa63e398e03f1cd39c97502ca783dc32786c39f05dafdef54155c38847542",
		91: "69b480d3a41488c56105f1e41141648f79fab067ab1379c36efb6686260a3da9",
		92: "150349a450a84d3ca36de71dc6c5fdd025ef93901ddab5eeb177e318405486b1",
		93: "b5a3e294dbc4b94556ff2cad7a75e6da2e53ac41244e305ae2512bf4d6b4f511",
		94: "598c06f72fc6c5bdb558b1a59fc53fe8143043dfaf7f708f5e220c46aad3cab5",
		95: "2bf02ea6e3e1f50b3853294183bcd51566ea1ef741e94d064abf7693424ec739",
		96: "fbe403773744ab020f995a239889877cc1f438c6861e935b2d2fd5c760a9088a",
		97: "25281afef03a84fe4902a59132fe305051674f4715feb9bb8fb6609c055e7223",
		98: "4bc8abb1a95bbe73c850697cb7bd2fc59c8eaa50336a4334484c95a931c55038",
		99: "c348c95ff57c40fe57bfd97b671b7a5f90421f83ee321d984f6ddf89c45262ef",
	},
}
