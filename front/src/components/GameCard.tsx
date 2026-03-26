import { Link } from "react-router-dom";
import { Star, Clock } from "lucide-react";

interface GameCardProps {
  id: string;
  title: string;
  image: string;
  rating?: number;
  hoursPlayed?: number;
  genre?: string;
  showRating?: boolean;
}

const GameCard = ({ id, title, image, rating, hoursPlayed, genre, showRating = true }: GameCardProps) => {
  return (
    <Link
      to={`/game/${id}`}
      className="group relative rounded-xl overflow-hidden bg-card border border-border/50 hover:border-primary/30 transition-all duration-300 hover:shadow-[0_8px_30px_-12px_hsl(var(--primary)/0.2)]"
    >
      <div className="aspect-[3/4] overflow-hidden">
        <img
          src={image}
          alt={title}
          className="w-full h-full object-cover transition-transform duration-500 group-hover:scale-105"
        />
        <div className="absolute inset-0 bg-gradient-to-t from-card via-card/20 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300" />
      </div>

      <div className="absolute bottom-0 left-0 right-0 p-3.5">
        <div className="glass rounded-lg p-3 translate-y-2 opacity-0 group-hover:translate-y-0 group-hover:opacity-100 transition-all duration-300 ease-out">
          <h3 className="font-semibold text-sm leading-tight line-clamp-1 mb-1.5">{title}</h3>
          <div className="flex items-center justify-between">
            {genre && (
              <span className="text-xs text-muted-foreground">{genre}</span>
            )}
            <div className="flex items-center gap-2 text-xs text-muted-foreground">
              {hoursPlayed !== undefined && (
                <span className="flex items-center gap-1">
                  <Clock className="w-3 h-3" />
                  {hoursPlayed}h
                </span>
              )}
              {showRating && rating !== undefined && (
                <span className="flex items-center gap-1 text-star-filled">
                  <Star className="w-3 h-3 fill-current" />
                  {rating.toFixed(1)}
                </span>
              )}
            </div>
          </div>
        </div>
      </div>

      {/* Always-visible title overlay */}
      <div className="absolute bottom-0 left-0 right-0 p-3 bg-gradient-to-t from-card/90 to-transparent group-hover:opacity-0 transition-opacity duration-300">
        <h3 className="font-medium text-sm line-clamp-1">{title}</h3>
      </div>
    </Link>
  );
};

export default GameCard;
