import { Link } from "react-router-dom";
import { ArrowRight, Gamepad2, Star, TrendingUp } from "lucide-react";
import Navbar from "@/components/Navbar";
import ReviewCard from "@/components/ReviewCard";
import GameCard from "@/components/GameCard";
import { games, reviews } from "@/lib/mockData";

const Home = () => {
  return (
    <div className="min-h-screen bg-background">
      <Navbar />

      {/* Hero */}
      <section className="relative pt-16 overflow-hidden">
        <div className="hero-gradient absolute inset-0" />
        <div className="container relative flex flex-col items-center text-center pt-24 pb-20 md:pt-32 md:pb-28">
          <div className="fade-up flex items-center gap-2 px-4 py-1.5 rounded-full glass-subtle text-xs font-medium text-muted-foreground mb-8">
            <TrendingUp className="w-3.5 h-3.5 text-primary" />
            Mais de 12.000 reviews publicadas
          </div>

          <h1 className="fade-up stagger-1 text-4xl sm:text-5xl md:text-6xl font-bold leading-[1.08] max-w-3xl mb-6">
            Sua biblioteca.{" "}
            <span className="text-gradient">Suas opiniões.</span>
          </h1>

          <p className="fade-up stagger-2 text-lg text-muted-foreground max-w-lg mb-10 leading-relaxed">
            Conecte sua conta Steam, avalie seus jogos e descubra o que a comunidade está jogando.
          </p>

          <div className="fade-up stagger-3 flex flex-col sm:flex-row gap-3">
            <Link
              to="/dashboard"
              className="flex items-center justify-center gap-2.5 px-7 py-3.5 rounded-xl bg-primary text-primary-foreground font-semibold text-sm hover:bg-primary/90 active:scale-[0.97] transition-all duration-200 glow-sm"
            >
              <Gamepad2 className="w-4 h-4" />
              Entrar com Steam
            </Link>
            <Link
              to="/dashboard"
              className="flex items-center justify-center gap-2 px-7 py-3.5 rounded-xl glass text-sm font-medium hover:bg-secondary transition-all duration-200 active:scale-[0.97]"
            >
              Explorar reviews
              <ArrowRight className="w-4 h-4" />
            </Link>
          </div>

          {/* Stats */}
          <div className="fade-up stagger-4 flex items-center gap-8 sm:gap-12 mt-16 text-center">
            {[
              { value: "12.4k", label: "Reviews" },
              { value: "3.2k", label: "Jogadores" },
              { value: "8.7k", label: "Jogos" },
            ].map((stat) => (
              <div key={stat.label}>
                <div className="text-2xl font-bold">{stat.value}</div>
                <div className="text-xs text-muted-foreground mt-0.5">{stat.label}</div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Popular Games */}
      <section className="container pb-20">
        <div className="flex items-center justify-between mb-8">
          <div>
            <h2 className="text-2xl font-bold">Em destaque</h2>
            <p className="text-sm text-muted-foreground mt-1">Jogos mais avaliados pela comunidade</p>
          </div>
          <Link
            to="/dashboard"
            className="text-sm text-primary hover:text-primary/80 font-medium flex items-center gap-1 transition-colors"
          >
            Ver todos <ArrowRight className="w-3.5 h-3.5" />
          </Link>
        </div>
        <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-4">
          {games.slice(0, 4).map((game, i) => (
            <div key={game.id} className={`fade-up stagger-${i + 1}`}>
              <GameCard {...game} />
            </div>
          ))}
        </div>
      </section>

      {/* Recent Reviews */}
      <section className="container pb-24">
        <div className="flex items-center justify-between mb-8">
          <div>
            <h2 className="text-2xl font-bold">Reviews recentes</h2>
            <p className="text-sm text-muted-foreground mt-1">O que a comunidade está dizendo</p>
          </div>
        </div>
        <div className="grid md:grid-cols-2 gap-4">
          {reviews.map((review, i) => (
            <div key={i} className={`fade-up stagger-${i + 1}`}>
              <ReviewCard {...review} />
            </div>
          ))}
        </div>
      </section>

      {/* Footer */}
      <footer className="border-t border-border/50">
        <div className="container py-8 flex flex-col sm:flex-row items-center justify-between gap-4 text-sm text-muted-foreground">
          <div className="flex items-center gap-2">
            <Gamepad2 className="w-4 h-4 text-primary" />
            <span className="font-medium text-foreground">GameVault</span>
          </div>
          <p>© 2026 GameVault. Não afiliado à Valve ou Steam.</p>
        </div>
      </footer>
    </div>
  );
};

export default Home;
