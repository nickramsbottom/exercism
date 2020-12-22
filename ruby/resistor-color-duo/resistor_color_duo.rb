# frozen_string_literal: true

# Class to hold static methods for interpreting resistor colors
class ResistorColorDuo
  ALL_COLORS = %w[black brown red orange yellow green blue violet grey white].freeze

  def self.value(colors)
    tens, ones = colors
    ALL_COLORS.find_index(ones) + ALL_COLORS.find_index(tens) * 10
  end
end
