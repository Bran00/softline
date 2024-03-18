"use client"
import styles from "./page.module.css";
import { useState } from 'react';
import { useAuth } from '../hooks/auth';
import Link from 'next/link';

export default function Home() {
  const { signIn, user, signOut } = useAuth(); 
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleLogin = () => {
    signIn({ username, password });
  }

  const handleLogout = () => {
    signOut();
  }

  return (
    <div>
      {user ? (
        <div className={styles.container} >
          <Link href="/listaProdutos">Lista de Produtos</Link>
          <Link href="/listaClientes">Lista de Clientes</Link>
          <button onClick={handleLogout}>
            Sair
          </button>
        </div>
      ) : (
        <div className={styles.login}>
          <h1>Login</h1>
          <input
            type="text"
            placeholder="Usuário"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
          <input
            type="password"
            placeholder="Senha"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <button onClick={handleLogin}>Entrar</button>
        </div>
      )}
    </div>
  );
};
