import express from 'express';

const app = express();

app.use(express.urlencoded({ extended: true }));
app.use(express.json());

app.post('/', (req, res) => {
  const { coupon, isValid } = req.body;

  if (isValid === 'valid') {
    const Discount = coupon.replace(/\D/g, '');

    return res.json({ Status: isValid, Discount });
  }

  return res.json({ Status: isValid, Discount: '' });
});

app.listen(9093);
