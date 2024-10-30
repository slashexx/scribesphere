import Header from "./Header";
import Hero from "./Hero";
import FeaturedPosts from "./FeaturedPosts";
import Footer from "./Footer";

export default function Home() {
    return (
      <div className="min-h-screen bg-white">
        <Header />
        <Hero />
        <FeaturedPosts />
        <Footer />
      </div>
    );
  }