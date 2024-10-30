import React from 'react';
import { TrendingUp, Clock } from 'lucide-react';

const featuredPosts = [
  {
    title: 'The Future of Web Development in 2024',
    excerpt: 'Exploring the latest trends and technologies shaping the web development landscape...',
    image: 'https://images.unsplash.com/photo-1498050108023-c5249f4df085?auto=format&fit=crop&q=80',
    readTime: '5 min',
    category: 'Technology',
  },
  {
    title: 'Sustainable Living: Small Changes, Big Impact',
    excerpt: 'Discover practical ways to reduce your environmental footprint while saving money...',
    image: 'https://images.unsplash.com/photo-1542601906990-b4d3fb778b09?auto=format&fit=crop&q=80',
    readTime: '4 min',
    category: 'Lifestyle',
  },
  {
    title: 'Hidden Gems: Southeast Asia\'s Best Kept Secrets',
    excerpt: 'Venture off the beaten path to discover authentic experiences and breathtaking locations...',
    image: 'https://images.unsplash.com/photo-1552733407-5d5c46c3bb3b?auto=format&fit=crop&q=80',
    readTime: '6 min',
    category: 'Travel',
  },
];

export default function FeaturedPosts() {
  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
      <div className="flex items-center justify-between mb-12">
        <div className="flex items-center space-x-2">
          <TrendingUp className="h-6 w-6 text-teal-600" />
          <h2 className="text-2xl font-bold text-slate-800">Featured Posts</h2>
        </div>
        <button className="text-teal-600 hover:text-teal-700 font-medium">
          View all posts
        </button>
      </div>

      <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
        {featuredPosts.map((post, index) => (
          <div 
            key={index} 
            className="group bg-white rounded-2xl shadow-sm overflow-hidden hover:shadow-md transition-all duration-300 border border-slate-100"
          >
            <div className="relative overflow-hidden">
              <img 
                src={post.image} 
                alt={post.title} 
                className="h-48 w-full object-cover transition-transform duration-300 group-hover:scale-105" 
              />
              <div className="absolute inset-0 bg-gradient-to-t from-black/20 to-transparent"></div>
            </div>
            <div className="p-6">
              <div className="flex items-center space-x-2 mb-4">
                <span className="text-sm text-teal-600 font-medium">{post.category}</span>
                <span className="text-slate-300">•</span>
                <span className="text-sm text-slate-500 flex items-center">
                  <Clock className="h-4 w-4 mr-1" />
                  {post.readTime}
                </span>
              </div>
              <h3 className="text-xl font-bold text-slate-800 mb-2 group-hover:text-teal-600 transition-colors">
                {post.title}
              </h3>
              <p className="text-slate-600 mb-4 line-clamp-2">{post.excerpt}</p>
              <button className="text-teal-600 font-medium group-hover:text-teal-700 transition-colors">
                Read more →
              </button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}