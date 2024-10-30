import { BookOpen, Search } from 'lucide-react';

export default function Header() {
  return (
    <header className="bg-white/70 backdrop-blur-md sticky top-0 z-50 border-b border-slate-100">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
        <div className="flex items-center justify-between">
          <div className="flex items-center space-x-2">
            <BookOpen className="h-8 w-8 text-teal-600" />
            <span className="text-2xl font-bold text-slate-800">ScribeSphere</span>
          </div>
          <div className="flex items-center space-x-6">
            <div className="relative hidden md:block">
              <input
                type="text"
                placeholder="Search articles..."
                className="w-64 pl-10 pr-4 py-2 rounded-full border border-slate-200 focus:ring-2 focus:ring-teal-500 focus:border-transparent bg-slate-50/50"
              />
              <Search className="absolute left-3 top-2.5 h-5 w-5 text-slate-400" />
            </div>
            <button className="bg-teal-600 text-white px-6 py-2 rounded-full hover:bg-teal-700 transition-all shadow-sm hover:shadow">
              Sign In
            </button>
          </div>
        </div>
      </div>
    </header>
  );
}