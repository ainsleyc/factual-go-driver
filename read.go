package factual

type ReadOpts struct {
  View string
  Filter *Filter
}

func (c Client) ReadPath(opts ReadOpts) (string, error) {
  path := "/t/" + opts.View

  if opts.Filter != nil {
    filter, err := opts.Filter.MarshalJSON()
    if err != nil {
      return "", err
    }
    path += "?filters=" + string(filter)
  }
  return path, nil
}

func (c Client) Read(opts ReadOpts) ([]byte, error) {
  return nil, nil
}
