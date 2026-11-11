┌──────────────────────────────────────┐
│ MAGIC NUMBER          8 bytes        │
├──────────────────────────────────────┤
│ FORMAT VERSION        2 bytes        │
├──────────────────────────────────────┤
│ NUMBER OF STORES      4 bytes        │
├──────────────────────────────────────┤
│                                      │
│ STORE 1                              │
│   Store ID Length     4 bytes        │
│   Store ID            Variable       │
│   Store Name Length   4 bytes        │
│   Store Name          Variable       │
│   Number of Entries   4 bytes        │
│                                      │
│   ENTRY 1                            │
│     Key Length        4 bytes        │
│     Key               Variable       │
│     Value Length      4 bytes        │
│     Value             Variable       │
│                                      │
│   ENTRY 2                            │
│     Key Length        4 bytes        │
│     Key               Variable       │
│     Value Length      4 bytes        │
│     Value             Variable       │
│                                      │
├──────────────────────────────────────┤
│ STORE 2                              │
│   Store ID Length     4 bytes        │
│   Store ID            Variable       │
│   Store Name Length   4 bytes        │
│   Store Name          Variable       │
│   Number of Entries   4 bytes        │
│   ...                                │
└──────────────────────────────────────┘