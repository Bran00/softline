'use client'
import axios from 'axios';

export const api = axios.create({
  baseUrl: process.env.BACKNODE,
})