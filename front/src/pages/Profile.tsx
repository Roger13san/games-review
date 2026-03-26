import { Calendar, Gamepad2, BookOpen } from "lucide-react";
import { Link } from "react-router-dom";
import Navbar from "@/components/Navbar";
import ReviewCard from "@/components/ReviewCard";
import GameCard from "@/components/GameCard";
import { games, reviews, userProfile } from "@/lib/mockData";

const Profile = () => {
  return (
    <div className="min-h-screen bg-background">
      <Navbar />
      <div className="container pt-24 pb-16">
        {/* Profile Header */}
        <div className="fade-up glass rounded-2xl p-6 sm:p-8 mb-10">
          <div className="flex flex-col sm:flex-row items-center sm:items-start gap-6">
            <img
              src={userProfile.avatar}
              alt={userProfile.name}
              className="w-24 h-24 rounded-2xl ring-2 ring-primary/30"
            />
            <div className="text-center sm:text-left flex-1">
              <h1 className="text-2xl font-bold">{userProfile.name}</h1>
              <p className="text-sm text-muted-foreground mt-1">
                Steam ID: {userProfile.steamId}
              </p>
              <div className="flex flex-wrap items-center justify-center sm:justify-start gap-5 mt-4 text-sm text-muted-foreground">
                <span className="flex items-center gap-1.5">
                  <Gamepad2 className="w-4 h-4 text-primary" />
                  <span className="text-foreground font-semibold">{userProfile.totalGames}</span> jogos
                </span>
                <span className="flex items-center gap-1.5">
                  <BookOpen className="w-4 h-4 text-primary" />
                  <span className="text-foreground font-semibold">{userProfile.gamesReviewed}</span> reviews
                </span>
                <span className="flex items-center gap-1.5">
                  <Calendar className="w-4 h-4" />
                  Membro desde {userProfile.memberSince}
                </span>
              </div>
            </div>
          </div>
        </div>

        {/* Recent Reviews */}
        <section className="fade-up stagger-1 mb-12">
          <h2 className="text-xl font-bold mb-6">Reviews recentes</h2>
          <div className="grid md:grid-cols-2 gap-4">
            {reviews.slice(0, 2).map((review, i) => (
              <ReviewCard key={i} {...review} />
            ))}
          </div>
        </section>

        {/* Favorite Games */}
        <section className="fade-up stagger-2">
          <h2 className="text-xl font-bold mb-6">Jogos avaliados</h2>
          <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-4">
            {games.slice(0, 5).map((game) => (
              <GameCard key={game.id} {...game} />
            ))}
          </div>
        </section>
      </div>
    </div>
  );
};

export default Profile;
