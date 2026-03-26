import { Link, useLocation } from "react-router-dom";
import { Gamepad2, User, Search } from "lucide-react";

const Navbar = () => {
  const location = useLocation();
  const isLoggedIn = location.pathname !== "/";

  return (
    <nav className="fixed top-0 left-0 right-0 z-50 glass">
      <div className="container flex items-center justify-between h-16">
        <Link to="/" className="flex items-center gap-2.5 group">
          <div className="w-8 h-8 rounded-lg bg-primary/20 flex items-center justify-center group-hover:glow-sm transition-shadow duration-300">
            <Gamepad2 className="w-4.5 h-4.5 text-primary" />
          </div>
          <span className="text-lg font-semibold tracking-tight">
            Game<span className="text-gradient">Vault</span>
          </span>
        </Link>

        <div className="hidden md:flex items-center gap-1">
          {[
            { to: "/", label: "Início" },
            { to: "/dashboard", label: "Biblioteca" },
            { to: "/profile", label: "Perfil" },
          ].map((link) => (
            <Link
              key={link.to}
              to={link.to}
              className={`px-3.5 py-2 rounded-lg text-sm font-medium transition-colors duration-200 ${
                location.pathname === link.to
                  ? "text-primary bg-primary/10"
                  : "text-muted-foreground hover:text-foreground hover:bg-secondary"
              }`}
            >
              {link.label}
            </Link>
          ))}
        </div>

        <div className="flex items-center gap-3">
          <button className="w-9 h-9 rounded-lg flex items-center justify-center text-muted-foreground hover:text-foreground hover:bg-secondary transition-colors duration-200 active:scale-95">
            <Search className="w-4 h-4" />
          </button>
          {isLoggedIn ? (
            <Link to="/profile" className="flex items-center gap-2 group">
              <img
                src="https://i.pravatar.cc/32?img=12"
                alt="Avatar"
                className="w-8 h-8 rounded-lg ring-2 ring-transparent group-hover:ring-primary/50 transition-all duration-200"
              />
            </Link>
          ) : (
            <Link
              to="/dashboard"
              className="flex items-center gap-2 px-4 py-2 rounded-lg bg-primary text-primary-foreground text-sm font-medium hover:bg-primary/90 active:scale-[0.97] transition-all duration-200"
            >
              <User className="w-3.5 h-3.5" />
              Entrar
            </Link>
          )}
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
