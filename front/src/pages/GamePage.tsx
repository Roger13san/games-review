import { useParams, Link } from "react-router-dom";
import { Star, Clock, Tag, ArrowLeft, PenLine } from "lucide-react";
import Navbar from "@/components/Navbar";
import ReviewCard from "@/components/ReviewCard";
import { StarRating } from "@/components/ReviewCard";
import { games, reviews } from "@/lib/mockData";

const GamePage = () => {
  const { id } = useParams();
  const game = games.find((g) => g.id === id) || games[0];
  const gameReviews = reviews.filter((r) => r.gameId === game.id);

  return (
    <div className="min-h-screen bg-background">
      <Navbar />

      {/* Banner */}
      <div className="relative h-64 sm:h-80 overflow-hidden">
        <img
          src={game.image}
          alt={game.title}
          className="w-full h-full object-cover"
        />
        <div className="absolute inset-0 bg-gradient-to-t from-background via-background/60 to-transparent" />
      </div>

      <div className="container -mt-20 relative pb-16">
        <Link
          to="/dashboard"
          className="fade-up inline-flex items-center gap-1.5 text-sm text-muted-foreground hover:text-foreground transition-colors mb-6"
        >
          <ArrowLeft className="w-3.5 h-3.5" />
          Voltar à biblioteca
        </Link>

        <div className="fade-up stagger-1">
          <h1 className="text-3xl sm:text-4xl font-bold mb-3">{game.title}</h1>
          <div className="flex flex-wrap items-center gap-4 text-sm text-muted-foreground mb-8">
            <span className="flex items-center gap-1.5">
              <Star className="w-4 h-4 text-star-filled fill-star-filled" />
              <span className="text-foreground font-semibold">{game.rating}</span>
            </span>
            <span className="flex items-center gap-1.5">
              <Clock className="w-4 h-4" />
              {game.hoursPlayed}h jogadas
            </span>
            <span className="flex items-center gap-1.5">
              <Tag className="w-4 h-4" />
              {game.genre}
            </span>
          </div>
        </div>

        {/* Action */}
        <div className="fade-up stagger-2 mb-12">
          <Link
            to="/create-review"
            className="inline-flex items-center gap-2 px-6 py-3 rounded-xl bg-primary text-primary-foreground text-sm font-semibold hover:bg-primary/90 active:scale-[0.97] transition-all duration-200 glow-sm"
          >
            <PenLine className="w-4 h-4" />
            Escrever review
          </Link>
        </div>

        {/* Reviews */}
        <div className="fade-up stagger-3">
          <h2 className="text-xl font-bold mb-6">
            Reviews da comunidade
            <span className="text-muted-foreground font-normal text-base ml-2">
              ({gameReviews.length})
            </span>
          </h2>
          {gameReviews.length > 0 ? (
            <div className="space-y-4 max-w-2xl">
              {gameReviews.map((review, i) => (
                <ReviewCard key={i} {...review} />
              ))}
            </div>
          ) : (
            <div className="glass rounded-xl p-8 text-center max-w-md">
              <p className="text-muted-foreground">Nenhuma review ainda.</p>
              <p className="text-sm text-muted-foreground/70 mt-1">Seja o primeiro a avaliar!</p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default GamePage;
