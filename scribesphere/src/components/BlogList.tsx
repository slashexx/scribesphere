import React, { useState } from 'react';
import { Search, Filter, BookOpen } from 'lucide-react';

// This will be replaced with Firebase data later
const categories = ['All', 'Technology', 'Lifestyle', 'Travel', 'Food', 'Health', 'Business'];

export default function BlogList() {
  const [selectedCategory, setSelectedCategory] = useState('All');
  const [searchQuery, setSearchQuery] = useState('');

  return (
    <div className="min-h-screen bg-slate-50/50">
      {/* Hero Section */}
      <div className="bg-gradient-to-br from-slate-900 via-slate-800 to-teal-900 text-white py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex items-center space-x-3 mb-6">
            <BookOpen className="h-8 w-8" />
            <h1 className="text-3xl font-bold">Explore Articles</h1>
          </div>
          <p className="text-slate-300 max-w-2xl">
            Discover thought-provoking articles across various topics. From technology to lifestyle,
            our writers bring you the best content worth your time.
          </p>
        </div>
      </div>

      {/* Filters and Search */}
      <div className="sticky top-16 z-40 bg-white/80 backdrop-blur-md border-b border-slate-100 py-4">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex flex-col md:flex-row md:items-center justify-between gap-4">
            <div className="flex items-center space-x-2 overflow-x-auto pb-2 md:pb-0 scrollbar-hide">
              <Filter className="h-5 w-5 text-slate-400 flex-shrink-0" />
              {categories.map((category) => (
                <button
                  key={category}
                  onClick={() => setSelectedCategory(category)}
                  className={`px-4 py-1.5 rounded-full text-sm font-medium whitespace-nowrap transition-all
                    ${selectedCategory === category
                      ? 'bg-teal-600 text-white'
                      : 'bg-slate-100 text-slate-600 hover:bg-slate-200'
                    }`}
                >
                  {category}
                </button>
              ))}
            </div>
            <div className="relative">
              <input
                type="text"
                placeholder="Search articles..."
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                className="w-full md:w-64 pl-10 pr-4 py-2 rounded-full border border-slate-200 focus:ring-2 focus:ring-teal-500 focus:border-transparent"
              />
              <Search className="absolute left-3 top-2.5 h-5 w-5 text-slate-400" />
            </div>
          </div>
        </div>
      </div>

      {/* Blog Grid */}
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
          {/* This is a placeholder card that will be mapped over Firebase data */}
          {[1, 2, 3, 4, 5, 6].map((_, index) => (
            <article
              key={index}
              className="group bg-white rounded-2xl shadow-sm overflow-hidden hover:shadow-md transition-all duration-300 border border-slate-100"
            >
              <div className="relative h-48 overflow-hidden">
                <img
                  src={`https://images.unsplash.com/photo-${[
                    '1498050108023-c5249f4df085',
                    '1542601906990-b4d3fb778b09',
                    '1552733407-5d5c46c3bb3b',
                    '1515378960530-7c0da6231fb1',
                    '1526374965328-7f61d4dc18c5',
                    '1504639725646-df3b0af6fb2c'
                  ][index]}?auto=format&fit=crop&q=80`}
                  alt="Blog cover"
                  className="w-full h-full object-cover transition-transform duration-300 group-hover:scale-105"
                />
                <div className="absolute inset-0 bg-gradient-to-t from-black/20 to-transparent"></div>
                <span className="absolute bottom-4 left-4 bg-white/90 backdrop-blur-md text-sm font-medium px-3 py-1 rounded-full text-teal-600">
                  {['Technology', 'Lifestyle', 'Travel', 'Food', 'Health', 'Business'][index]}
                </span>
              </div>
              <div className="p-6">
                <div className="flex items-center space-x-4 mb-4">
                  <img
                    src={`https://i.pravatar.cc/150?img=${index + 1}`}
                    alt="Author"
                    className="h-10 w-10 rounded-full"
                  />
                  <div>
                    <h3 className="text-sm font-medium text-slate-700">John Doe</h3>
                    <p className="text-sm text-slate-500">Mar 16, 2024 · 5 min read</p>
                  </div>
                </div>
                <h2 className="text-xl font-bold text-slate-800 mb-2 group-hover:text-teal-600 transition-colors line-clamp-2">
                  {[
                    'The Future of AI in Software Development',
                    'Mindful Living in a Digital Age',
                    'Hidden Gems of Southeast Asia',
                    'Modern Architecture Trends',
                    'Sustainable Tech Solutions',
                    'Remote Work Best Practices'
                  ][index]}
                </h2>
                <p className="text-slate-600 mb-4 line-clamp-2">
                  Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor
                  incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam...
                </p>
                <button className="text-teal-600 font-medium group-hover:text-teal-700 transition-colors inline-flex items-center">
                  Read more <span className="ml-1">→</span>
                </button>
              </div>
            </article>
          ))}
        </div>

        {/* Load More Button */}
        <div className="flex justify-center mt-12">
          <button className="bg-white text-teal-600 font-medium px-8 py-3 rounded-full border-2 border-teal-600 hover:bg-teal-50 transition-colors">
            Load More Articles
          </button>
        </div>
      </div>
    </div>
  );
}