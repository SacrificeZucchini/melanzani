module Tone
  C = 0
  CS = 1
  D = 2
  DS = 3
  E = 4
  F = 5
  FS = 6
  G = 7
  GS = 8
  A = 9
  AS = 10
  B = 11
end

class Note
  def initialize(tone, octave)
    @tone = tone
    @octave = octave
  end

  def from_offset(offset)
    n = @tone + @octave + offset
    tone = n % 12
    octave = n / 12
    Note::new(tone, octave)
  end

  def to_s()
    tone = Tone.constants.find do |t|
      Tone.const_get(t) == @tone
    end
    "#{tone} #{@octave}"
  end
end

puts "Enter note..."

input = ""

while input != "X" && input != "x" do
  input = gets

  tone, offset = input.split(" ")

  tone = Tone.const_get(tone.to_sym)
  offset = offset.to_i

  note = Note::new(tone, 0)
  puts "> #{note.from_offset(offset)}"
  puts
end
