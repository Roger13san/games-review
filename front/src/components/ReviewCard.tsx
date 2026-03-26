import { Star } from "lucide-react";
import { Link } from "react-router-dom";

interface ReviewCardProps {
  gameTitle: string;
  gameImage: string;
  gameId: string;
  author: string;
  authorAvatar: string;
  rating: number;
  content: string;
  date: string;
}

const StarRating = ({ rating, max = 5 }: { rating: number; max?: number }) => (
  <div className="flex gap-0.5">
    {Array.from({ length: max }, (_, i) => (
      <Star
        key={i}
        className={`w-3.5 h-3.5 ${
          i < rating ? "text-star-filled fill-star-filled" : "text-star-empty"
        }`}
      />
    ))}
  </div>
);

const ReviewCard = ({
  gameTitle,
  gameImage,
  gameId,
  author,
  authorAvatar,
  rating,
  content,
  date,
}: ReviewCardProps) => {
  return (
    <div className="glass rounded-xl p-4 hover:border-primary/20 transition-all duration-300 group">
      <div className="flex gap-3.5">
        <Link to={`/game/${gameId}`} className="shrink-0">
          <img
            src={gameImage}
            alt={gameTitle}
            className="w-14 h-20 rounded-lg object-cover ring-1 ring-border/50 group-hover:ring-primary/30 transition-all duration-300"
          />
        </Link>
        <div className="flex-1 min-w-0">
          <div className="flex items-start justify-between gap-2 mb-1.5">
            <div>
              <Link
                to={`/game/${gameId}`}
                className="font-semibold text-sm hover:text-primary transition-colors line-clamp-1"
              >
                {gameTitle}
              </Link>
              <div className="flex items-center gap-2 mt-0.5">
                <img
                  src={authorAvatar}
                  alt={author}
                  className="w-4 h-4 rounded-full"
                />
                <span className="text-xs text-muted-foreground">{author}</span>
                <span className="text-xs text-muted-foreground/50">·</span>
                <span className="text-xs text-muted-foreground/70">{date}</span>
              </div>
            </div>
            <StarRating rating={rating} />
          </div>
          <p className="text-sm text-secondary-foreground/80 line-clamp-2 leading-relaxed">
            {content}
          </p>
        </div>
      </div>
    </div>
  );
};

export { StarRating };
export default ReviewCard;
