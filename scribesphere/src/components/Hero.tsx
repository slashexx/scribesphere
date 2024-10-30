import React from 'react';
import { ChevronRight } from 'lucide-react';

export default function Hero() {
  return (
    <div className="relative bg-gradient-to-br from-slate-900 via-slate-800 to-teal-900 text-white py-32">
      <div className="absolute inset-0 bg-[url('https://images.unsplash.com/photo-1519681393784-d120267933ba')] opacity-10 bg-cover bg-center"></div>
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 relative">
        <div className="max-w-3xl">
          <h1 className="text-5xl font-bold mb-6 bg-gradient-to-r from-white to-teal-200 bg-clip-text text-transparent">
            Discover Stories That Matter
          </h1>
          <p className="text-xl text-slate-300 mb-8 leading-relaxed">
            Join our community of writers and readers. Share your thoughts, learn from others,
            and stay inspired with the latest stories across various topics.
          </p>
          <button className="bg-white/10 backdrop-blur-md text-white px-8 py-3 rounded-full font-medium hover:bg-white/20 transition-all inline-flex items-center border border-white/20">
            Start Reading <ChevronRight className="ml-2 h-5 w-5" />
          </button>
        </div>
      </div>
    </div>
  );
}