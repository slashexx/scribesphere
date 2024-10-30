import React from 'react';
import { Mail } from 'lucide-react';

const categories = [
  { name: 'Technology', count: 24, color: 'bg-teal-50 hover:bg-teal-100' },
  { name: 'Lifestyle', count: 18, color: 'bg-teal-50 hover:bg-teal-100' },
  { name: 'Travel', count: 12, color: 'bg-teal-50 hover:bg-teal-100' },
  { name: 'Food', count: 9, color: 'bg-teal-50 hover:bg-teal-100' },
];

export default function Footer() {
  return (
    <div className="bg-slate-50/50 py-16 border-t border-slate-100">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="grid md:grid-cols-2 gap-16">
          <div>
            <h2 className="text-2xl font-bold text-slate-800 mb-6">Popular Categories</h2>
            <div className="grid grid-cols-2 gap-4">
              {categories.map((category, index) => (
                <div
                  key={index}
                  className={`${category.color} rounded-xl p-4 transition-all cursor-pointer border border-slate-100`}
                >
                  <h3 className="font-semibold text-slate-800">{category.name}</h3>
                  <p className="text-slate-600">{category.count} articles</p>
                </div>
              ))}
            </div>
          </div>

          <div className="bg-white rounded-2xl p-8 shadow-sm border border-slate-100">
            <div className="flex items-center space-x-3 mb-4">
              <Mail className="h-6 w-6 text-teal-600" />
              <h2 className="text-2xl font-bold text-slate-800">Newsletter</h2>
            </div>
            <p className="text-slate-600 mb-6">
              Get the latest posts delivered right to your inbox. No spam, unsubscribe anytime.
            </p>
            <form className="flex space-x-2">
              <input
                type="email"
                placeholder="Enter your email"
                className="flex-1 px-4 py-2 rounded-full border border-slate-200 focus:ring-2 focus:ring-teal-500 focus:border-transparent bg-slate-50"
              />
              <button className="bg-teal-600 text-white px-6 py-2 rounded-full hover:bg-teal-700 transition-all shadow-sm hover:shadow">
                Subscribe
              </button>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
}