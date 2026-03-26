import { useState } from "react";
import { Star, ArrowLeft, Send } from "lucide-react";
import { Link } from "react-router-dom";
import Navbar from "@/components/Navbar";
import { games } from "@/lib/mockData";

const CreateReview = () => {
  const [rating, setRating] = useState(0);
  const [hoveredRating, setHoveredRating] = useState(0);
  const [content, setContent] = useState("");
  const game = games[0]; // Default to first game for demo

  const displayRating = hoveredRating || rating;

  return (
    <div className="min-h-screen bg-background">
      <Navbar />
      <div className="container max-w-2xl pt-24 pb-16">
        <Link
          to={`/game/${game.id}`}
          className="fade-up inline-flex items-center gap-1.5 text-sm text-muted-foreground hover:text-foreground transition-colors mb-8"
        >
          <ArrowLeft className="w-3.5 h-3.5" />
          Voltar ao jogo
        </Link>

        <div className="fade-up stagger-1 glass rounded-2xl p-6 sm:p-8">
          {/* Game info */}
          <div className="flex items-center gap-4 mb-8 pb-6 border-b border-border/50">
            <img
              src={game.image}
              alt={game.title}
              className="w-16 h-10 rounded-lg object-cover ring-1 ring-border/50"
            />
            <div>
              <h2 className="font-semibold">{game.title}</h2>
              <p className="text-xs text-muted-foreground">{game.genre}</p>
            </div>
          </div>

          {/* Rating */}
          <div className="mb-8">
            <label className="text-sm font-medium text-muted-foreground mb-3 block">
              Sua nota
            </label>
            <div className="flex gap-1.5">
              {[1, 2, 3, 4, 5].map((star) => (
                <button
                  key={star}
                  onMouseEnter={() => setHoveredRating(star)}
                  onMouseLeave={() => setHoveredRating(0)}
                  onClick={() => setRating(star)}
                  className="p-1 transition-transform duration-150 hover:scale-110 active:scale-95"
                >
                  <Star
                    className={`w-8 h-8 transition-colors duration-150 ${
                      star <= displayRating
                        ? "text-star-filled fill-star-filled"
                        : "text-star-empty"
                    }`}
                  />
                </button>
              ))}
            </div>
            {rating > 0 && (
              <p className="text-xs text-muted-foreground mt-2">
                {rating === 1 && "Fraco"}
                {rating === 2 && "Regular"}
                {rating === 3 && "Bom"}
                {rating === 4 && "Ótimo"}
                {rating === 5 && "Obra-prima"}
              </p>
            )}
          </div>

          {/* Text */}
          <div className="mb-8">
            <label className="text-sm font-medium text-muted-foreground mb-3 block">
              Sua review
            </label>
            <textarea
              value={content}
              onChange={(e) => setContent(e.target.value)}
              placeholder="O que você achou desse jogo?"
              rows={6}
              className="w-full px-4 py-3 rounded-xl bg-secondary border border-border/50 text-sm placeholder:text-muted-foreground resize-none focus:outline-none focus:ring-2 focus:ring-primary/50 focus:border-primary/50 transition-all duration-200 leading-relaxed"
            />
            <p className="text-xs text-muted-foreground mt-2 text-right">
              {content.length}/2000
            </p>
          </div>

          {/* Submit */}
          <button
            disabled={rating === 0 || content.length < 10}
            className="w-full flex items-center justify-center gap-2 px-6 py-3.5 rounded-xl bg-primary text-primary-foreground font-semibold text-sm hover:bg-primary/90 active:scale-[0.97] transition-all duration-200 disabled:opacity-40 disabled:pointer-events-none glow-sm"
          >
            <Send className="w-4 h-4" />
            Publicar review
          </button>
        </div>
      </div>
    </div>
  );
};

export default CreateReview;
