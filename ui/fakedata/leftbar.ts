import { Channel } from "@/models/Channel";
// import { Stream } from "@/models/Stream";
import { User } from "@/models/User";

const users: User[] = [
    {
        id: 1,
        avatar: "",
        bio: "This is my bio",
        username: "MrBeast",
        email: "beast123@gmail.com",
        phoneNumber: "",
        password: "",
        birth: new Date(),
    },
];

// const streams: Stream[] = [
//     {
//         id: 1,
//         ownerId: 1,
//         viewers: 345,
//         title: "Ace on day 27",
//         chats: [],
//         startedTime: new Date(),
//         tags: ["English", "Troll"],
//         category: "Troll VietNam",
//     },
//     {
//         id: 2,
//         ownerId: 2,
//         viewers: 500,
//         title: "Island Hopping Adventure",
//         chats: [],
//         startedTime: new Date(),
//         tags: ["Travel", "Adventure"],
//         category: "Travel Adventures",
//     },
//     {
//         id: 3,
//         ownerId: 3,
//         viewers: 120,
//         title: "Coding Marathon",
//         chats: [],
//         startedTime: new Date(),
//         tags: ["Technology", "Coding"],
//         category: "Tech Talks",
//     },
//     {
//         id: 4,
//         ownerId: 4,
//         viewers: 1000,
//         title: "Epic Gaming Night",
//         chats: [],
//         startedTime: new Date(),
//         tags: ["Gaming", "Entertainment"],
//         category: "Gaming Fun",
//     },
//     {
//         id: 5,
//         ownerId: 5,
//         viewers: 345,
//         title: "Ace on day 27",
//         chats: [],
//         startedTime: new Date(),
//         tags: ["English", "Troll"],
//         category: "Troll VietNam",
//     },
//     {
//         id: 6,
//         ownerId: 6,
//         viewers: 500,
//         title: "Island Hopping Adventure",
//         chats: [],
//         startedTime: new Date(),
//         tags: ["Travel", "Adventure"],
//         category: "Travel Adventures",
//     },
//     {
//         id: 7,
//         ownerId: 7,
//         viewers: 120,
//         title: "Coding Marathon",
//         chats: [],
//         startedTime: new Date(),
//         tags: ["Technology", "Coding"],
//         category: "Tech Talks",
//     },
//     {
//         id: 8,
//         ownerId: 8,
//         viewers: 1000,
//         title: "Epic Gaming Night",
//         chats: [],
//         startedTime: new Date(),
//         tags: ["Gaming", "Entertainment"],
//         category: "Gaming Fun",
//     },
//     {
//         id: 9,
//         ownerId: 9,
//         viewers: 1000,
//         title: "Epic Gaming Night",
//         chats: [],
//         startedTime: new Date(),
//         tags: ["Gaming", "Entertainment"],
//         category: "Gaming Fun",
//     },
//     {
//         id: 10,
//         ownerId: 10,
//         viewers: 345,
//         title: "Ace on day 27",
//         chats: [],
//         startedTime: new Date(),
//         tags: ["English", "Troll"],
//         category: "Troll VietNam",
//     },
//     {
//         id: 11,
//         ownerId: 11,
//         viewers: 500,
//         title: "Island Hopping Adventure",
//         chats: [],
//         startedTime: new Date(),
//         tags: ["Travel", "Adventure"],
//         category: "Travel Adventures",
//     },
//     {
//         id: 12,
//         ownerId: 12,
//         viewers: 120,
//         title: "Coding Marathon",
//         chats: [],
//         startedTime: new Date(),
//         tags: ["Technology", "Coding"],
//         category: "Tech Talks",
//     },
//     {
//         id: 13,
//         ownerId: 13,
//         viewers: 1000,
//         title: "Epic Gaming Night",
//         chats: [],
//         startedTime: new Date(),
//         tags: ["Gaming", "Entertainment"],
//         category: "Gaming Fun",
//     },
// ];

const channels: Channel[] = [];

export { users, channels };
