import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';

const useAuthCheck = () => {
  const [isLoading, setIsLoading] = useState(true);
  const router = useRouter();

  useEffect(() => {
    const verifyAuth = async () => {
      const response = await fetch('http://localhost:8080/verify', { credentials: 'include' });
      
      if (!response.ok) {
        router.push('/login');
      }
      setIsLoading(false);
    };

    verifyAuth();
  }, [router]);

  return isLoading;
};

export default useAuthCheck;
