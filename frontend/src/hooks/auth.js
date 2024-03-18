"use client"
import { createContext, useContext, useState, useEffect } from 'react';
import { api } from '../services/api';
import { useRouter } from 'next/navigation';

export const AuthContext = createContext();

const AuthProvider = ({ children, onSignIn, onSignOut }) => {
  const router = useRouter()
  const [user, setUser] = useState(null);
  const [token, setToken] = useState(null);

  async function signIn({ username, password }) {
    try {
      const response = await api.post(`http://localhost:8080/login`, { username, password });
      const { user, token } = response.data;

      localStorage.setItem('user', JSON.stringify(user));
      localStorage.setItem('token', token);

      setUser(user);
      setToken(token);
      api.defaults.headers.common['Authorization'] = `Bearer ${token}`;

      if (onSignIn) {
        onSignIn();
      }
    } catch (error) {
      if (error.response) {
        alert("Senha ou Usuário errados!");
      }
    }
  }

  function signOut() {
    localStorage.removeItem('user');
    localStorage.removeItem('token');
    setUser(null);
    setToken(null);
    api.defaults.headers.common['Authorization'] = '';

    if (onSignOut) {
      onSignOut();
    }
  }

  useEffect(() => {
    const storedUser = localStorage.getItem('user');
    const storedToken = localStorage.getItem('token');
    if (storedUser && storedToken) {
      setUser(JSON.parse(storedUser));
      setToken(storedToken);
      api.defaults.headers.common['Authorization'] = `Bearer ${storedToken}`;
    } else {
      router.push('/');
    }
  }, []);

  return (
    <AuthContext.Provider
      value={{
        user,
        token,
        signIn,
        signOut,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};

const useAuth = () => {
  const context = useContext(AuthContext);
  return context;
};

export { AuthProvider, useAuth };
