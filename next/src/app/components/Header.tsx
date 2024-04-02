import React from 'react';
import Link from 'next/link';
import { useRouter } from 'next/navigation';

const Header: React.FC = () => {
  const router = useRouter();

  const handleLogout = async () => {
    const response = await fetch('http://localhost:8080/logout', {
      credentials: 'include',
    })
    if (response.ok) {
      router.push('/login');
    } else {
      alert('ログアウトに失敗しました。')
    }
    const data = await response.json();
    console.log(data);
  }

  return (
    <header className="bg-gray-800 text-white px-4 py-2 flex justify-between items-center">
      <div>
        <Link href="/users" className="text-xl font-bold">
          ホーム
        </Link>
      </div>
      <div className="flex items-center">
        <Link href="/mypage" className="text-white mr-8 hover:underline">
          マイページ
        </Link>
        <button className="px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600" onClick={handleLogout}>
          ログアウト
        </button>
      </div>
    </header>
  );
};

export default Header;