"use client"
import { useState } from 'react';
import Link from 'next/link';
import axios from 'axios';
import styles from './CadastroProduto.module.css';

const CadastroProduto = () => {
  const [nome, setNome] = useState('');
  const [descricao, setDescricao] = useState('');
  const [codigoBarras, setCodigoBarras] = useState('');
  const [valorVenda, setValorVenda] = useState('');
  const [pesoBruto, setPesoBruto] = useState('');
  const [pesoLiquido, setPesoLiquido] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const token = localStorage.getItem('token');
      const response = await axios.post(
        'http://localhost:8080/produtos',
        {
          nome: nome,
          descricao: descricao,
          codigo_barras: codigoBarras,
          valor_venda: parseFloat(valorVenda),
          peso_bruto: parseFloat(pesoBruto),
          peso_liquido: parseFloat(pesoLiquido)
        },
        {
          headers: {
            Authorization: `Bearer ${token}`
          }
        }
      );

      setNome('');
      setDescricao('');
      setCodigoBarras('');
      setValorVenda('');
      setPesoBruto('');
      setPesoLiquido('');

      alert('Produto criado com sucesso!');
    } catch (error) {
      alert('Falha ao criar o produto. Por favor, tente novamente.');
    }
  };

  return (
    <div className={styles.container}>
      <h1>Cadastro de Produto</h1>
      <form onSubmit={handleSubmit}>
        <input type="text" placeholder="Nome" value={nome} onChange={(e) => setNome(e.target.value)} />
        <input type="text" placeholder="Descrição" value={descricao} onChange={(e) => setDescricao(e.target.value)} />
        <input
          type="text"
          placeholder="Código de Barras"
          value={codigoBarras}
          onChange={(e) => setCodigoBarras(e.target.value)}
        />
        <input
          type="number"
          placeholder="Valor de Venda"
          value={valorVenda}
          onChange={(e) => setValorVenda(e.target.value)}
        />
        <input
          type="number"
          placeholder="Peso Bruto em KG"
          value={pesoBruto}
          onChange={(e) => setPesoBruto(e.target.value)}
        />
        <input
          type="number"
          placeholder="Peso Líquido em KG"
          value={pesoLiquido}
          onChange={(e) => setPesoLiquido(e.target.value)}
        />
        <button type="submit">Cadastrar</button>

        <Link href="/listaProdutos">Lista de Produtos</Link>
      </form>
    </div>
  );
};

export default CadastroProduto
