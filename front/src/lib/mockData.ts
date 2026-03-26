export const games = [
  { id: "1", title: "Elden Ring", image: "https://cdn.cloudflare.steamstatic.com/steam/apps/1245620/header.jpg", rating: 4.8, hoursPlayed: 142, genre: "Action RPG" },
  { id: "2", title: "Hollow Knight", image: "https://cdn.cloudflare.steamstatic.com/steam/apps/367520/header.jpg", rating: 4.9, hoursPlayed: 67, genre: "Metroidvania" },
  { id: "3", title: "Hades", image: "https://cdn.cloudflare.steamstatic.com/steam/apps/1145360/header.jpg", rating: 4.7, hoursPlayed: 88, genre: "Roguelike" },
  { id: "4", title: "Celeste", image: "https://cdn.cloudflare.steamstatic.com/steam/apps/504230/header.jpg", rating: 4.9, hoursPlayed: 34, genre: "Platformer" },
  { id: "5", title: "Disco Elysium", image: "https://cdn.cloudflare.steamstatic.com/steam/apps/632470/header.jpg", rating: 4.6, hoursPlayed: 52, genre: "RPG" },
  { id: "6", title: "Stardew Valley", image: "https://cdn.cloudflare.steamstatic.com/steam/apps/413150/header.jpg", rating: 4.8, hoursPlayed: 210, genre: "Simulation" },
  { id: "7", title: "Dead Cells", image: "https://cdn.cloudflare.steamstatic.com/steam/apps/588650/header.jpg", rating: 4.5, hoursPlayed: 45, genre: "Roguelike" },
  { id: "8", title: "Sekiro", image: "https://cdn.cloudflare.steamstatic.com/steam/apps/814380/header.jpg", rating: 4.7, hoursPlayed: 78, genre: "Action" },
];

export const reviews = [
  {
    gameTitle: "Elden Ring",
    gameImage: "https://cdn.cloudflare.steamstatic.com/steam/apps/1245620/header.jpg",
    gameId: "1",
    author: "Lucas Ribeiro",
    authorAvatar: "https://i.pravatar.cc/32?img=12",
    rating: 5,
    content: "Uma obra-prima absoluta. O mundo aberto mais bem feito que já joguei. Cada canto tem algo para descobrir, cada boss é memorável.",
    date: "2 dias atrás",
  },
  {
    gameTitle: "Hollow Knight",
    gameImage: "https://cdn.cloudflare.steamstatic.com/steam/apps/367520/header.jpg",
    gameId: "2",
    author: "Marina Costa",
    authorAvatar: "https://i.pravatar.cc/32?img=5",
    rating: 5,
    content: "A exploração nesse jogo é viciante. Cada nova área traz surpresas. O combate é simples mas muito satisfatório.",
    date: "5 dias atrás",
  },
  {
    gameTitle: "Hades",
    gameImage: "https://cdn.cloudflare.steamstatic.com/steam/apps/1145360/header.jpg",
    gameId: "3",
    author: "Rafael Mendes",
    authorAvatar: "https://i.pravatar.cc/32?img=8",
    rating: 4,
    content: "Gameplay excelente e narrativa que se integra perfeitamente ao loop roguelike. Nunca me cansei de jogar.",
    date: "1 semana atrás",
  },
  {
    gameTitle: "Celeste",
    gameImage: "https://cdn.cloudflare.steamstatic.com/steam/apps/504230/header.jpg",
    gameId: "4",
    author: "Ana Beatriz",
    authorAvatar: "https://i.pravatar.cc/32?img=9",
    rating: 5,
    content: "Mais do que um platformer difícil, é uma história sobre superação. A trilha sonora é de outro mundo.",
    date: "2 semanas atrás",
  },
];

export const userProfile = {
  name: "Lucas Ribeiro",
  avatar: "https://i.pravatar.cc/128?img=12",
  steamId: "76561198012345678",
  gamesReviewed: 24,
  totalGames: 147,
  memberSince: "Mar 2023",
};
