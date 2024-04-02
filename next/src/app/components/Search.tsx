"use client";

import React, { ChangeEvent, useState } from 'react';

interface SearchInputProps {
  onSearch: (searchTerm: string) => void;
}

const Search: React.FC<SearchInputProps> = ({ onSearch }) => {
  const [searchTerm, setSearchTerm] = useState('');

  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    const newSearchTerm = event.target.value;
    setSearchTerm(newSearchTerm)
    onSearch(newSearchTerm);
  };

  return (
    <div className="mb-4 flex flex-row items-center">
      <input
        type="text"
        placeholder="ユーザーを検索"
        value={searchTerm}
        onChange={handleChange}
        className="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 mr-2"
      />
    </div>
  );
};

export default Search;